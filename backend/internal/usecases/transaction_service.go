package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/gateways"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *entities.Transaction) (string, error)
	FindByID(ctx context.Context, id string) (*entities.Transaction, error)
}

type TransactionQueue interface {
	PublishTransaction(ctx context.Context, tx *entities.Transaction) error
}

type TransactionService struct {
	repo    TransactionRepository
	queue   TransactionQueue
	convert gateways.CurrencyGateway
}

func NewTransactionService(repo TransactionRepository, queue TransactionQueue, convert gateways.CurrencyGateway) *TransactionService {
	return &TransactionService{
		repo:    repo,
		queue:   queue,
		convert: convert,
	}
}

// Cria uma transação, valida, e publica na fila para persistência assíncrona
func (s *TransactionService) CreateTransaction(ctx context.Context, description string, date time.Time, valueUSD float64) (*entities.Transaction, error) {
	tx, err := entities.NewTransaction(description, date, valueUSD)
	if err != nil {
		return nil, err
	}

	if err := s.queue.PublishTransaction(ctx, tx); err != nil {
		return nil, err
	}

	return tx, nil
}

// Recupera e converte a transação para moeda desejada
func (s *TransactionService) GetTransactionConverted(ctx context.Context, id string, currency string) (*gateways.CurrencyConversion, error) {
	tx, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tx == nil {
		return nil, errors.New("transação não encontrada")
	}

	conversion, err := s.convert.ConvertUSDTo(currency, tx.Date, tx.ValueUSD)
	if err != nil {
		return nil, err
	}

	conversion.FromCurrency = "USD"
	conversion.ToCurrency = currency

	return conversion, nil
}
