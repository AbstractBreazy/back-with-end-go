package models

import (
	"back-with-end-go/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Connection struct {
	*gorm.DB
}

func ConnectToDB(config config.Config) (*Connection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true",
		config.DatabaseConfiguration.User,
		config.DatabaseConfiguration.Password,
		config.DatabaseConfiguration.Server,
		config.DatabaseConfiguration.Port,
		config.DatabaseConfiguration.Database)
	print(dsn)
	db, err := gorm.Open("postgres", "host="+
		config.DatabaseConfiguration.Server+" user="+config.DatabaseConfiguration.User+
		" dbname="+config.DatabaseConfiguration.Database+" sslmode=disable password="+config.DatabaseConfiguration.Password)
	if err != nil {
		print(err)
		return nil, err
	}
	print("connected!")
	return &Connection{db}, err
}
