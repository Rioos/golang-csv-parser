package utils

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// GetZeroValueFromNull returns empty string if s == "NULL" if s != "NULL" returns s
func GetZeroValueFromNull(s string) string {
	if s == "NULL" {
		return ""
	}
	return s
}

// GetBoolFromString returns s as boolean equivalent
func GetBoolFromString(s string) bool {
	result, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// GetFloat32FromString returns s to float32 equivalent
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

// GetTimeFromString returns s to Time equivalent using format "2006-01-02"
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
