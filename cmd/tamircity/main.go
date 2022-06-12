package main

import (
	"log"

	"github.com/joho/godotenv"
	dbModels "github.com/mustafakocatepe/Tamircity/pkg/models/db"
	postgres "github.com/mustafakocatepe/Tamircity/pkg/store/shared/db"
)

func main() {
	//Set enviroment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	db.AutoMigrate(&dbModels.Brand{}, &dbModels.Model{}, &dbModels.FixType{}, &dbModels.DeviceType{})
	log.Println("Migrations done")
}
