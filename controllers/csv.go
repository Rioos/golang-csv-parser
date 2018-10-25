package controllers

import (
	"fmt"
	"golang-csv-parser/services"
	"log"
	"net/http"
	"time"
)

// HandleRequestCSV handles request to '/csv' route
func HandleRequestCSV(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprint(w, "Couldn't find file. Did you sent it as a form-data and on 'file' field?")
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "text/plain", "text/csv":
		services.ReadCSV(file, w)
		break
	default:
		jsonResponse(w, http.StatusBadRequest, "The file mime type is not valid. Accept only plain text or CSV")
	}

	elapsed := time.Since(start)
	log.Printf("Request took %s", elapsed)
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
