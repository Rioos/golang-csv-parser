package utils

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func GetBoolFromString(s string) bool {
	result, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetFloat32FromString(s string) float32 {
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

func GetTimeFromString(s string) time.Time {
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
