package main

import (
	"fmt"
	"golang-csv-parser/models"
	"golang-csv-parser/routes"
	"log"

	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable  port=5432 user=postgres dbname=csv_neoway password=postgres")
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Client{})
	defer db.Close()
	fmt.Println("Database ready")
	routes.Listen()
	fmt.Println("Listening")
}
