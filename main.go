package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"golang-csv-parser/models"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

func main() {
	start := time.Now()
	var txn = createTxn()

	stmt, err := txn.Prepare(pq.CopyIn("clients", "cpf", "last_purchase_store", "most_frequent_store", "private", "incomplete", "last_purchase", "medium_purchase_value", "last_pruchase_value"))
	if err != nil {
		log.Fatal(err)
	}

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

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
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

func createTxn() *sql.Tx {
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
	var private = getBoolFromString(values[1])
	var incomplete = getBoolFromString(values[2])
	var lastPurchase = getTimeFromString(values[3])
	var mediumPurchaseValue = getFloat32FromString(values[4])
	var lastPruchaseValue = getFloat32FromString(values[5])
	client := models.Client{
		CPF:                 values[0],
		Private:             private,
		Incomplete:          incomplete,
		LastPurchase:        lastPurchase,
		MediumPurchaseValue: mediumPurchaseValue,
		LastPruchaseValue:   lastPruchaseValue,
		MostFrequentStore:   values[6],
		LastPurchaseStore:   values[7]}
	return client, nil
}

func getBoolFromString(s string) bool {
	result, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getFloat32FromString(s string) float32 {
	if s != "NULL" {
		s = strings.Replace(s, ",", ".", 1)
		result, err := strconv.ParseFloat(s, 32)
		if err != nil {
			log.Fatal(err)
		}
		return float32(result)
	}
	return 0
}

func getTimeFromString(s string) time.Time {
	var result time.Time
	if s != "NULL" {
		_lastPurchase, err := time.Parse("2006-01-02", s)
		if err != nil {
			log.Fatal(err)
		}
		result = _lastPurchase
	}
	return result
}

// func usingReadAll() {
// 	start := time.Now()
// 	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=csv_neoway sslmode=disable password=postgres")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db.CreateTable(&Client{})
// 	defer db.Close()
// 	csvFile, _ := os.Open("base_teste.txt")
// 	csvReader := csv.NewReader(bufio.NewReader(csvFile))
// 	csvReader.Comma = ' '
// 	csvReader.TrimLeadingSpace = true
// 	values, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(values)
// 	elapsed := time.Since(start)
// 	log.Printf("Binomial took %s", elapsed)
// }
