package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/neglarken/date_name_parser/config"
	"github.com/neglarken/date_name_parser/internal/requests"
)

var (
	configPath string = "config/config.yaml"
)

func init() {
	if err := godotenv.Load(); err != nil {
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

	_, err = reqs.Auth("admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
	res, err := reqs.GetDataFromDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
}
