package controllers

import (
	"fmt"
	"golang-csv-parser/services"
	"log"
	"net/http"
	"time"
)

func HandleRequestCSV(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", "Couldn't find file. Did you sent it as a form-data and on 'file' field?")
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "text/plain":
		services.ReadCSV(file)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	case "text/csv":
		services.ReadCSV(file)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	default:
		jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
	}
	elapsed := time.Since(start)
	log.Printf("Request took %s", elapsed)
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
