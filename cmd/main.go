package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/neglarken/date_name_parser/config"
	"github.com/neglarken/date_name_parser/internal/requests"
)

var (
	configPath string = "config/config.yaml"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	reqs, err := requests.NewRequests(config)
	if err != nil {
		log.Fatal(err)
	}

	_, err = reqs.Auth(config.Login, config.Password)
	if err != nil {
		log.Fatal(err)
	}
	res, err := reqs.GetDataFromDB(config.SqlLimit)
	if err != nil {
		log.Fatal(err)
	}
	err = reqs.PostDataToMySQL(res)
	if err != nil {
		log.Fatal(err)
	}
}
