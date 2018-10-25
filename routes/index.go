package routes

import (
	"fmt"
	"golang-csv-parser/controllers"
	"net/http"
)

func Listen() {
	http.HandleFunc("/csv", controllers.HandleRequestCSV)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
