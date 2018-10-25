package routes

import (
	"fmt"
	"golang-csv-parser/controllers"
	"net/http"
)

// Listen listens on port 8080
func Listen() {
	http.HandleFunc("/csv", controllers.HandleRequestCSV)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on 8080")
	if err != nil {
		fmt.Println(err)
	}
}
