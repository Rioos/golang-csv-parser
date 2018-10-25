package controllers

import (
	"fmt"
	"golang-csv-parser/services"
	"net/http"
)

func HandleRequestCSV(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "text/plain":
		services.ReadCSV(file)
	case "text/csv":
		services.ReadCSV(file)
	default:
		jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
	}
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
