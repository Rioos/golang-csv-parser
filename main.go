package main

import (
	"fmt"
	"golang-csv-parser/models"
	"golang-csv-parser/routes"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func main() {
	var connString = fmt.Sprintf(
		"host=db sslmode=disable user=%s dbname=%s password=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Client{})
	defer db.Close()
	fmt.Println("Database ready")
	routes.Listen()
	fmt.Println("Listening")
}
