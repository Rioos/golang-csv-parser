package routes

import (
	"fmt"
	"golang-csv-parser/controllers"
	"net/http"
)

// Listen listens on port 8080
func Listen() {
	http.HandleFunc("/csv", controllers.HandleRequestCSV)
	fmt.Println("Listening")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
