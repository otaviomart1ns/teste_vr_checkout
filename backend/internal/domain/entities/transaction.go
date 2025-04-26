package entities

import (
	"errors"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/pkg/utils"
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
	if !utils.IsAlphanumeric(description) {
		return nil, errors.New("descrição inválida: apenas caracteres alfanuméricos são permitidos")
	}
	if valueUSD <= 0 {
		return nil, errors.New("valor da transação deve ser positivo")
	}

	return &Transaction{
		Description: description,
		Date:        date,
		ValueUSD:    utils.RoundToCents(valueUSD),
	}, nil
}