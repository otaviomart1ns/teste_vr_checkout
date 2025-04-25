package entities

import (
	"errors"
	"time"
)

type Transaction struct {
	ID          string
	Description string
	Date        time.Time
	ValueUSD    float64
}

func NewTransaction(description string, date time.Time, valueUSD float64) (*Transaction, error) {
	if len(description) == 0 || len(description) > 50 {
		return nil, errors.New("descrição inválida: deve ter entre 1 e 50 caracteres")
	}
	if !isAlphanumeric(description) {
		return nil, errors.New("descrição inválida: apenas caracteres alfanuméricos são permitidos")
	}
	if valueUSD <= 0 {
		return nil, errors.New("valor da transação deve ser positivo")
	}

	return &Transaction{
		Description: description,
		Date:        date,
		ValueUSD:    roundToCents(valueUSD),
	}, nil
}

func roundToCents(value float64) float64 {
	return float64(int(value*100+0.5)) / 100
}

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !(r >= 'a' && r <= 'z') &&
			!(r >= 'A' && r <= 'Z') &&
			!(r >= '0' && r <= '9') &&
			!(r == ' ') {
			return false
		}
	}
	return true
}
