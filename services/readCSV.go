package services

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"golang-csv-parser/models"
	"golang-csv-parser/utils"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/lib/pq"
)

// ReadCSV reads a CSV file and persist it on database returns number of lines read
func ReadCSV(file multipart.File, w http.ResponseWriter) {
	var db = createConnection()
	var tx = createTx(db)
	var stmt = startCopyStmt(tx)
	var csvReader = createReader(file)
	var lineCount int
	for {
		if lineCount > 0 {
			client, err := readNextLine(csvReader)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			stmt.Exec(
				client.CPF,
				client.LastPurchaseStore,
				client.MostFrequentStore,
				client.Private,
				client.Incomplete,
				client.LastPurchase,
				client.MediumPurchaseValue,
				client.LastPruchaseValue,
				client.HasValidCPF,
				client.HasValidLastPurchaseStore,
				client.HasValidMostFrequentStore)
		}
		lineCount++
	}
	_, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v %v", "Read a total of", lineCount)
}

func createReader(file multipart.File) *csv.Reader {
	csvReader := csv.NewReader(file)
	csvReader.Comma = ' '
	csvReader.TrimLeadingSpace = true
	csvReader.ReuseRecord = true
	return csvReader
}

func createConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=csv_neoway sslmode=disable password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createTx(db *sql.DB) *sql.Tx {
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return txn
}

func readNextLine(r *csv.Reader) (models.Client, error) {
	values, err := r.Read()
	if err == io.EOF {
		return models.Client{}, err
	} else if err != nil {
		log.Fatal(err)
	}
	client := models.Client{
		CPF:                 utils.RemoveNonAlphanumeric(utils.GetZeroValueFromNull(values[0])),
		Private:             utils.GetBoolFromString(values[1]),
		Incomplete:          utils.GetBoolFromString(values[2]),
		LastPurchase:        utils.GetTimeFromString(values[3]),
		MediumPurchaseValue: utils.GetFloat32FromString(values[4]),
		LastPruchaseValue:   utils.GetFloat32FromString(values[5]),
		MostFrequentStore:   utils.RemoveNonAlphanumeric(utils.GetZeroValueFromNull(values[6])),
		LastPurchaseStore:   utils.RemoveNonAlphanumeric(utils.GetZeroValueFromNull(values[7]))}
	client.HasValidCPF = client.ValidateCPF()
	client.HasValidLastPurchaseStore = client.ValidateLastPurchaseStore()
	client.HasValidMostFrequentStore = client.ValidateMostFrequentStore()
	return client, nil
}

func startCopyStmt(tx *sql.Tx) *sql.Stmt {
	stmt, err := tx.Prepare(
		pq.CopyIn(
			"clients",
			"cpf",
			"last_purchase_store",
			"most_frequent_store",
			"private",
			"incomplete",
			"last_purchase",
			"medium_purchase_value",
			"last_pruchase_value",
			"has_valid_cpf",
			"has_valid_last_purchase_store",
			"has_valid_most_frequent_store"))
	if err != nil {
		log.Fatal(err)
	}
	return stmt

}
