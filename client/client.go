package client

import (
	"time"

	"github.com/Nhanderu/brdoc"
)

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

func (c Client) ValidateCPF() bool {
	return brdoc.IsCPF(c.CPF)
}

func (c Client) ValidateLastPurchaseStore() bool {
	return brdoc.IsCNPJ(c.LastPurchaseStore)
}

func (c Client) ValidateMostFrequentStore() bool {
	return brdoc.IsCNPJ(c.MostFrequentStore)
}
