package entities_test

import (
	"testing"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction_Success(t *testing.T) {
	tx, err := entities.NewTransaction("Compra Teste", time.Now(), 99.99)

	assert.NoError(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, "Compra Teste", tx.Description)
	assert.Equal(t, 99.99, tx.ValueUSD)
}

func TestNewTransaction_EmptyDescription(t *testing.T) {
	tx, err := entities.NewTransaction("", time.Now(), 50.0)

	assert.Nil(t, tx)
	assert.EqualError(t, err, "descrição inválida: deve ter entre 1 e 50 caracteres")
}

func TestNewTransaction_TooLongDescription(t *testing.T) {
	desc := "Lorem ipsum dolor sit amet, consectetur adipiscing elit 123" // > 50 caracteres
	tx, err := entities.NewTransaction(desc, time.Now(), 10.0)

	assert.Nil(t, tx)
	assert.EqualError(t, err, "descrição inválida: deve ter entre 1 e 50 caracteres")
}

func TestNewTransaction_InvalidCharacters(t *testing.T) {
	tx, err := entities.NewTransaction("Compra @@@", time.Now(), 100.0)

	assert.Nil(t, tx)
	assert.EqualError(t, err, "descrição inválida: apenas caracteres alfanuméricos são permitidos")
}

func TestNewTransaction_ZeroAmount(t *testing.T) {
	tx, err := entities.NewTransaction("Compra Zero", time.Now(), 0)

	assert.Nil(t, tx)
	assert.EqualError(t, err, "valor da transação deve ser positivo")
}

func TestNewTransaction_RoundsToCents(t *testing.T) {
	tx, err := entities.NewTransaction("Compra Arredondada", time.Now(), 10.239)

	assert.NoError(t, err)
	assert.Equal(t, 10.24, tx.ValueUSD)
}
