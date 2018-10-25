package models

import (
	"time"

	"github.com/Nhanderu/brdoc"
)

// Client represents a row in the clients table
type Client struct {
	CPF                       string `gorm:"type:varchar(20);index"`
	LastPurchaseStore         string `gorm:"type:varchar(20);index"`
	MostFrequentStore         string `gorm:"type:varchar(20);index"`
	Private                   bool
	Incomplete                bool
	LastPurchase              time.Time `gorm:"index"`
	MediumPurchaseValue       float32   `gorm:"index"`
	LastPruchaseValue         float32   `gorm:"index"`
	HasValidCPF               bool
	HasValidLastPurchaseStore bool
	HasValidMostFrequentStore bool
}

// ValidateCPF returns true if Client CPF is empty string or is a valid CPF
func (c Client) ValidateCPF() bool {
	return brdoc.IsCPF(c.CPF)
}

// ValidateLastPurchaseStore returns true if Client LastPurchaseStore is empty string or is a valid CNPJ
func (c Client) ValidateLastPurchaseStore() bool {
	return brdoc.IsCNPJ(c.LastPurchaseStore)
}

// ValidateMostFrequentStore returns true if Client MostFrequentStore is empty string or is a valid CNPJ
func (c Client) ValidateMostFrequentStore() bool {
	return brdoc.IsCNPJ(c.MostFrequentStore)
}
