package models

import (
	"time"

	"github.com/Nhanderu/brdoc"
)

// Client represents a row in the clients table
type Client struct {
	CPF                 string
	LastPurchaseStore   string
	MostFrequentStore   string
	Private             bool
	Incomplete          bool
	LastPurchase        time.Time
	MediumPurchaseValue float32
	LastPruchaseValue   float32
}

// ValidateCPF returns true if Client CPF is empty string or is a valid CPF
func (c Client) ValidateCPF() bool {
	return trueIfZeroValue(c.CPF) || brdoc.IsCPF(c.CPF)
}

// ValidateLastPurchaseStore returns true if Client LastPurchaseStore is empty string or is a valid CNPJ
func (c Client) ValidateLastPurchaseStore() bool {
	return trueIfZeroValue(c.LastPurchaseStore) || brdoc.IsCNPJ(c.LastPurchaseStore)
}

// ValidateMostFrequentStore returns true if Client MostFrequentStore is empty string or is a valid CNPJ
func (c Client) ValidateMostFrequentStore() bool {
	return trueIfZeroValue(c.MostFrequentStore) || brdoc.IsCNPJ(c.MostFrequentStore)
}

// Returns true if string is empty
func trueIfZeroValue(s string) bool {
	if s == "" {
		return true
	}
	return false
}
