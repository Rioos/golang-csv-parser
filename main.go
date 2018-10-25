package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"golang-csv-parser/models"
	"golang-csv-parser/utils"
	"io"
	"log"
	"os"

	"github.com/lib/pq"
)

func main() {
	var tx = createTx()
	var stmt = startCopyStmt(tx)
	var csvReader = createReader()
	for {
		client, err := readNextLine(csvReader)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if client.ValidateCPF() && client.ValidateLastPurchaseStore() && client.ValidateMostFrequentStore() {
			stmt.Exec(client.CPF, client.LastPurchaseStore, client.MostFrequentStore, client.Private, client.Incomplete, client.LastPurchase, client.MediumPurchaseValue, client.LastPruchaseValue)
		}
	}
	_, err := stmt.Exec()
	err = stmt.Close()
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func createReader() *csv.Reader {
	csvFile, err := os.Open("base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	csvReader.Comma = ' '
	csvReader.TrimLeadingSpace = true
	return csvReader
}

func createTx() *sql.Tx {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=csv_neoway sslmode=disable password=postgres")
	if err != nil {
		log.Fatal(err)
	}
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
	var private = utils.GetBoolFromString(values[1])
	var incomplete = utils.GetBoolFromString(values[2])
	var lastPurchase = utils.GetTimeFromString(values[3])
	var mediumPurchaseValue = utils.GetFloat32FromString(values[4])
	var lastPruchaseValue = utils.GetFloat32FromString(values[5])
	client := models.Client{
		CPF:                 values[0],
		Private:             private,
		Incomplete:          incomplete,
		LastPurchase:        lastPurchase,
		MediumPurchaseValue: mediumPurchaseValue,
		LastPruchaseValue:   lastPruchaseValue,
		MostFrequentStore:   values[6],
		LastPurchaseStore:   values[7]}
	fmt.Println(client)
	return client, nil
}

func startCopyStmt(tx *sql.Tx) *sql.Stmt {
	stmt, err := tx.Prepare(pq.CopyIn("clients", "cpf", "last_purchase_store", "most_frequent_store", "private", "incomplete", "last_purchase", "medium_purchase_value", "last_pruchase_value"))
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}
