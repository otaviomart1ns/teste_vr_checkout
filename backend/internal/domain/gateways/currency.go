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
	// Converte um valor de USD para a moeda destino com base na data (ou datas anteriores até 6 meses)
	ConvertUSDTo(currency string, date time.Time, amount float64) (*CurrencyConversion, error)
}
