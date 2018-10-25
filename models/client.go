package models

import (
	"time"

	"github.com/Nhanderu/brdoc"
)

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

func (c Client) ValidateCPF() bool {
	return trueIfZeroValue(c.CPF) || brdoc.IsCPF(c.CPF)
}

func (c Client) ValidateLastPurchaseStore() bool {
	return trueIfZeroValue(c.LastPurchaseStore) || brdoc.IsCNPJ(c.LastPurchaseStore)
}

func (c Client) ValidateMostFrequentStore() bool {
	return trueIfZeroValue(c.MostFrequentStore) || brdoc.IsCNPJ(c.MostFrequentStore)
}

func trueIfZeroValue(s string) bool {
	if s == "" {
		return true
	}
	return false
}
