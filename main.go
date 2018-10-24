package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
  Client - A client that the database consume
*/
type Client struct {
	CPF                 string `gorm:"size:255;index"`
	LastPurchaseStore   string `gorm:"size:255;index"`
	MostFrequentStore   string `gorm:"size:255;index"`
	Private             bool
	Incomplete          bool
	LastPurchase        time.Time
	MediumPurchaseValue float32 `gorm:"index"`
	LastPruchaseValue   float32 `gorm:"index"`
}

func main() {
	start := time.Now()
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=csv_neoway sslmode=disable password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	db.CreateTable(&Client{})
	defer db.Close()
	csvFile, _ := os.Open("base_teste.txt")
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	csvReader.Comma = ' '
	csvReader.TrimLeadingSpace = true

	for {
		values, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		private, err := strconv.ParseBool(values[1])
		incomplete, err := strconv.ParseBool(values[2])
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
		client := Client{
			CPF:                 values[0],
			Private:             private,
			Incomplete:          incomplete,
			LastPurchase:        lastPurchase,
			MediumPurchaseValue: mediumPurchaseValue32,
			LastPruchaseValue:   lastPruchaseValue32,
			MostFrequentStore:   values[6],
			LastPurchaseStore:   values[7]}
		db.Create(&client)
	}

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func usingReadAll() {
	start := time.Now()
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=csv_neoway sslmode=disable password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	db.CreateTable(&Client{})
	defer db.Close()
	csvFile, _ := os.Open("base_teste.txt")
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	csvReader.Comma = ' '
	csvReader.TrimLeadingSpace = true
	values, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(values)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
