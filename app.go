package main

import (
	"back-with-end-go/config"
	"back-with-end-go/managers"
	"back-with-end-go/models"
	"back-with-end-go/response"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type Env struct {
	Configuration config.Config
	EntityManager managers.EntityManager
}

func main() {
	// Creates a global instance 'logger'
	logger := logrus.New()
	// JSONFormatter formats logs into parsable json
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}
	// Creates global instance variable based on Env struct
	env := &Env{}
	// Fetching database configuration data from passed JSON
	Configuration, err := config.GetConfig()
	if err != nil {
		panic(err.Error())
	}
	// initialize a new db connection
	connection, err := models.ConnectToDB(Configuration)
	if err != nil {
		panic(err.Error())
	}
	// migrates for given models
	connection.DB.AutoMigrate(
		&models.Entity{})
	connection.DB.LogMode(true)
	env.EntityManager = managers.InitEntityManager(connection)

	// Creates a router without any middleware
	r := chi.NewRouter()

	// Base middleware stack
	// RequestID is a middleware that injects a 'request id' into context
	// of each request.
	r.Use(middleware.RequestID)

	// RealIP is a middleware that sets a http.Request's RemoteAddr to the results
	// of parsing either the X-Forwarded-For header or the X-Real-IP header.
	r.Use(middleware.RealIP)

	// Logger Middleware writes a logs the start and end of each request,
	// along with useful data, what the response status was, and how long
	// long it took to return.
	r.Use(middleware.Logger)

	// Recoverer middleware 'recovers' from any panic and returns a 500 if there was one.
	r.Use(middleware.Recoverer)

	// SetContentType is a middleware that forces response Content-Type.
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Route("/api/status", func(r chi.Router) {
		r.Get("/check", env.CheckEntity)
		r.Post("/process", env.ProcessEntity)
	})
	// start listening :5432
	err = http.ListenAndServe(":8090", r)
	if err != nil {
		log.Print(err.Error())
		return
	}
}

func check(err error, w http.ResponseWriter, resp *response.StatusResponse, message string, code int) bool {
	if err != nil {
		resp.Status = false
		if len(message) == 0 {
			resp.ResponseMessage = err.Error()
		} else {
			resp.ResponseMessage = message + ": " + err.Error()
		}
		resp.ResponseInternalCode = code
		w.Write(resp.GetJSON())
		return true
	}
	return false
}
