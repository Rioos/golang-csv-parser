package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"golang-csv-parser/client"
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
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=csv_neoway sslmode=disable password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := txn.Prepare(pq.CopyIn("clients", "cpf", "last_purchase_store", "most_frequent_store", "private", "incomplete", "last_purchase", "medium_purchase_value", "last_pruchase_value"))

	if err != nil {
		log.Fatal(err)
	}

	csvFile, err := os.Open("base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	csvReader.Comma = ' '
	csvReader.TrimLeadingSpace = true
	var count = 0

	for {
		values, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		private, err := strconv.ParseBool(values[1])
		if err != nil {
			log.Fatal(err)
		}
		incomplete, err := strconv.ParseBool(values[2])
		if err != nil {
			log.Fatal(err)
		}
		var mediumPurchaseValue float32
		var lastPruchaseValue float32
		var lastPurchase time.Time

		if values[3] != "NULL" {
			_lastPurchase, err := time.Parse("2006-01-02", values[3])
			if err != nil {
				log.Fatal(err)
			}
			lastPurchase = _lastPurchase
		}

		if values[4] != "NULL" {
			values[4] = strings.Replace(values[4], ",", ".", 1)
			_mediumPurchaseValue, err := strconv.ParseFloat(values[4], 32)
			if err != nil {
				log.Fatal(err)
			}
			mediumPurchaseValue = float32(_mediumPurchaseValue)
		}

		if values[5] != "NULL" {
			values[5] = strings.Replace(values[5], ",", ".", 1)
			_lastPruchaseValue, err := strconv.ParseFloat(values[5], 32)
			if err != nil {
				log.Fatal(err)
			}
			lastPruchaseValue = float32(_lastPruchaseValue)
		}

		var lastPruchaseValue32 = float32(lastPruchaseValue)
		var mediumPurchaseValue32 = float32(mediumPurchaseValue)
		client := client.Client{
			CPF:                 values[0],
			Private:             private,
			Incomplete:          incomplete,
			LastPurchase:        lastPurchase,
			MediumPurchaseValue: mediumPurchaseValue32,
			LastPruchaseValue:   lastPruchaseValue32,
			MostFrequentStore:   values[6],
			LastPurchaseStore:   values[7]}
		if client.ValidateCPF() && client.ValidateLastPurchaseStore() && client.ValidateMostFrequentStore() {
			_, err = stmt.Exec(client.CPF, client.LastPurchaseStore, client.MostFrequentStore, client.Private, client.Incomplete, client.LastPurchase, client.MediumPurchaseValue, client.LastPruchaseValue)
		} else {
			count++
		}
	}

	fmt.Println(count)
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
