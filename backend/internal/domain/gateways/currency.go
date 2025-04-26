package gateways

import (
	"time"
)

type CurrencyConversion struct {
	FromCurrency string
	ToCurrency   string
	Rate         float64
	Converted    float64
	DateUsed     time.Time
}

type CurrencyGateway interface {
	ConvertUSDTo(currency string, date time.Time, amount float64) (*CurrencyConversion, error)
}

type CurrencyMetadataGateway interface {
	GetAvailableCurrencies() ([]string, error)
}