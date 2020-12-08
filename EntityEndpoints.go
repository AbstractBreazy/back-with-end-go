package main

import (
	"back-with-end-go/models"
	"back-with-end-go/response"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strconv"
)

func (env *Env) CheckEntity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response.New()
	fetchedResults, err := env.EntityManager.EntityDataStore.Get()
	if check(err, w, resp, "Error while clearing the table", 100) {
		return
	}
	content, err := json.Marshal(struct {
		ResultArray    *[]models.Entity         `json:"result_array"`
		StatusResponse *response.StatusResponse `json:"status_response"`
	}{fetchedResults, resp})
	if check(err, w, resp, "error marshalling response", 101) {
		return
	}
	w.Write(content)
	return
}

func (env *Env) ProcessEntity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response.New()
	value := r.FormValue("value")
	data := r.FormValue("data")
	typeOfRequest := r.FormValue("type")
	if len(typeOfRequest) <= 0 {
		if check(errors.New("undefined type request"), w, resp, "", 102) {
			return
		}
	}
	switch typeOfRequest {
	case "prepare":
		{
			if len(data) < 1 || len(value) < 1 {
				if check(errors.New("empty values"), w, resp, "", 103) {
					return
				}
			}
			if env.EntityManager.IsNumeric(data) == false {
				if check(errors.New("data isnt Numeric"), w, resp, "", 104) {
					return
				}
			}
			dataNum, _ := strconv.ParseFloat(data, 64)
			if math.Signbit(dataNum) == true {
				if check(errors.New("dataNum is negative"), w, resp, "", 105) {
					return
				}
			}
			newEntity, err := env.EntityManager.PostEntity(value, dataNum)
			if check(err, w, resp, "", 106) {
				return
			}
			content, err := json.Marshal(struct {
				StatusReponse *response.StatusResponse `json:"status_reponse"`
				Entity        *models.Entity           `json:"entity"`
			}{resp, newEntity})
			if check(err, w, resp, "error marshalling response", 107) {
				return
			}
			w.Write(content)
			return
		}
	case "start processing":
		{
			if len(data) < 1 || len(value) < 1 {
				if check(errors.New("empty values"), w, resp, "", 108) {
					return
				}
			}
			if env.EntityManager.IsNumeric(data) == false {
				if check(errors.New("data isnt Numeric"), w, resp, "", 109) {
					return
				}
			}
			dataNum, _ := strconv.ParseFloat(data, 64)
			if math.Signbit(dataNum) == true {
				if check(errors.New("dataNum is negative"), w, resp, "", 110) {
					return
				}
			}
			resultSlice, err := env.EntityManager.ProcessEntity(value, dataNum)
			if check(err, w, resp, "error marshalling response", 111) {
				return
			}
			content, err := json.Marshal(struct {
				StatusResponce *response.StatusResponse `json:"status_responce"`
				Results        []float64                `json:"results"`
			}{resp, resultSlice})
			if check(err, w, resp, "error marshalling response", 112) {
				return
			}
			w.Write(content)
			return
		}
	case "clear":
		{
			err := env.EntityManager.EntityDataStore.ClearTable()
			if check(err, w, resp, "error marshalling response", 113) {
				return
			}
			resp.ResponseMessage = "Entity Table has been cleared!"
			content, err := json.Marshal(struct {
				StatusResponce *response.StatusResponse `json:"status_responce"`
			}{resp})
			w.Write(content)
			return
		}
	default:
		{
			if check(errors.New("undefined type request"), w, resp, "", 114) {
				return
			}
		}
	}
}
