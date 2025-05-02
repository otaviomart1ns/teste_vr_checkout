package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/gateways"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct{ mock.Mock }

func (m *MockTransactionRepository) FindByID(ctx context.Context, id string) (*entities.Transaction, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetLatestTransactions(ctx context.Context, limit int32) ([]*entities.Transaction, error) {
	args := m.Called(ctx, limit)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) Save(ctx context.Context, tx *entities.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

type MockProducer struct{ mock.Mock }

func (m *MockProducer) PublishTransaction(ctx context.Context, tx *entities.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

type MockConverter struct{ mock.Mock }

func (m *MockConverter) ConvertUSDTo(desc string, date time.Time, amount float64) (*gateways.CurrencyConversion, error) {
	args := m.Called(desc, date, amount)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gateways.CurrencyConversion), args.Error(1)
}

func TestCreateTransaction_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	mockProducer := new(MockProducer)
	svc := usecases.NewTransactionService(mockProducer, mockRepo, nil)

	mockProducer.On("PublishTransaction", mock.Anything, mock.Anything).Return(nil)

	err := svc.CreateTransaction("Compra Teste", time.Now(), 100.0)
	assert.NoError(t, err)
	mockProducer.AssertCalled(t, "PublishTransaction", mock.Anything, mock.Anything)
}

func TestCreateTransaction_EmptyDescription(t *testing.T) {
	svc := usecases.NewTransactionService(nil, nil, nil)

	err := svc.CreateTransaction("", time.Now(), 100.0)
	assert.EqualError(t, err, "descrição inválida: deve ter entre 1 e 50 caracteres")
}

func TestCreateTransaction_TooLongDescription(t *testing.T) {
	svc := usecases.NewTransactionService(nil, nil, nil)

	desc := "Lorem ipsum dolor sit amet, consectetur adipiscing elit 123" // > 50 chars
	err := svc.CreateTransaction(desc, time.Now(), 100.0)
	assert.EqualError(t, err, "descrição inválida: deve ter entre 1 e 50 caracteres")
}

func TestCreateTransaction_InvalidCharacters(t *testing.T) {
	svc := usecases.NewTransactionService(nil, nil, nil)

	err := svc.CreateTransaction("Compra @@@", time.Now(), 100.0)
	assert.EqualError(t, err, "descrição inválida: apenas caracteres alfanuméricos e espaço são permitidos")
}

func TestCreateTransaction_ZeroAmount(t *testing.T) {
	svc := usecases.NewTransactionService(nil, nil, nil)

	err := svc.CreateTransaction("Compra Vazia", time.Now(), 0)
	assert.EqualError(t, err, "valor deve ser maior que 0 e no máximo 99.999,99")
}

func TestCreateTransaction_TooHighAmount(t *testing.T) {
	svc := usecases.NewTransactionService(nil, nil, nil)

	err := svc.CreateTransaction("Compra Carissima", time.Now(), 100000.00)
	assert.EqualError(t, err, "valor deve ser maior que 0 e no máximo 99.999,99")
}

func TestGetTransactionWithConversion_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	mockConverter := new(MockConverter)

	svc := usecases.NewTransactionService(nil, mockRepo, mockConverter)

	tx := &entities.Transaction{
		ID:          "123",
		Description: "Teste",
		Date:        time.Now(),
		ValueUSD:    100.0,
	}
	converted := &gateways.CurrencyConversion{
		ToCurrency: "Brazil-Real",
		Rate:       5.0,
		Converted:  500.0,
		DateUsed:   tx.Date,
	}

	mockRepo.On("FindByID", mock.Anything, "123").Return(tx, nil)
	mockConverter.On("ConvertUSDTo", "Brazil-Real", tx.Date, tx.ValueUSD).Return(converted, nil)

	resultTx, resultConv, err := svc.GetTransactionWithConversion("123", "Brazil-Real")

	assert.NoError(t, err)
	assert.Equal(t, tx, resultTx)
	assert.Equal(t, converted, resultConv)
}

func TestGetTransactionWithConversion_TransactionNotFound(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	svc := usecases.NewTransactionService(nil, mockRepo, nil)

	mockRepo.On("FindByID", mock.Anything, "123").Return(nil, nil)

	_, _, err := svc.GetTransactionWithConversion("123", "Brazil-Real")
	assert.EqualError(t, err, "transação não encontrada")
}

func TestGetTransactionWithConversion_RepoError(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	svc := usecases.NewTransactionService(nil, mockRepo, nil)

	mockRepo.On("FindByID", mock.Anything, "123").Return(nil, assert.AnError)

	_, _, err := svc.GetTransactionWithConversion("123", "Brazil-Real")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao buscar transação")
}

func TestGetTransactionWithConversion_ConversionError(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	mockConverter := new(MockConverter)

	svc := usecases.NewTransactionService(nil, mockRepo, mockConverter)

	tx := &entities.Transaction{
		ID:          "123",
		Description: "Teste",
		Date:        time.Now(),
		ValueUSD:    100.0,
	}

	mockRepo.On("FindByID", mock.Anything, "123").Return(tx, nil)
	mockConverter.On("ConvertUSDTo", "Brazil-Real", tx.Date, tx.ValueUSD).Return(nil, assert.AnError)

	_, _, err := svc.GetTransactionWithConversion("123", "Brazil-Real")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro na conversão")
}

func TestGetLatestTransactions_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	svc := usecases.NewTransactionService(nil, mockRepo, nil)

	txList := []*entities.Transaction{
		{ID: "1"}, {ID: "2"},
	}

	mockRepo.On("GetLatestTransactions", mock.Anything, int32(2)).Return(txList, nil)

	result, err := svc.GetLatestTransactions(2)
	assert.NoError(t, err)
	assert.Equal(t, txList, result)
}

func TestGetLatestTransactions_RepoError(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	svc := usecases.NewTransactionService(nil, mockRepo, nil)

	mockRepo.On("GetLatestTransactions", mock.Anything, int32(2)).Return(nil, assert.AnError)

	_, err := svc.GetLatestTransactions(2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao buscar últimas transações")
}
