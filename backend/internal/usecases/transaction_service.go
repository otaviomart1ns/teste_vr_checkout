package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/gateways"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/pkg/utils"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *entities.Transaction) error
	FindByID(ctx context.Context, id string) (*entities.Transaction, error)
	GetLatestTransactions(ctx context.Context, limit int32) ([]*entities.Transaction, error)
}

type TransactionService struct {
	producer  gateways.TransactionProducer
	repo      TransactionRepository
	converter gateways.CurrencyGateway
}

func NewTransactionService(
	producer gateways.TransactionProducer,
	repo TransactionRepository,
	converter gateways.CurrencyGateway,
) *TransactionService {
	return &TransactionService{
		producer:  producer,
		repo:      repo,
		converter: converter,
	}
}

func (s *TransactionService) CreateTransaction(description string, date time.Time, amount float64) error {
	if len(description) == 0 || len(description) > 50 {
		return fmt.Errorf("descrição inválida: deve ter entre 1 e 50 caracteres")
	}
	if !utils.IsAlphanumeric(description) {
		return fmt.Errorf("descrição inválida: apenas caracteres alfanuméricos e espaço são permitidos")
	}
	if amount <= 0 || amount > 99999.99 {
		return fmt.Errorf("valor deve ser maior que 0 e no máximo 99.999,99")
	}

	tx := &entities.Transaction{
		ID:          utils.GenerateUUID(),
		Description: description,
		Date:        date,
		ValueUSD:    amount,
	}

	ctx := context.Background()
	return s.producer.PublishTransaction(ctx, tx)
}

func (s *TransactionService) GetTransactionWithConversion(id, currency string) (*entities.Transaction, *gateways.CurrencyConversion, error) {
	ctx := context.Background()

	tx, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao buscar transação: %w", err)
	}
	if tx == nil {
		return nil, nil, fmt.Errorf("transação não encontrada")
	}

	converted, err := s.converter.ConvertUSDTo(currency, tx.Date, tx.ValueUSD)
	if err != nil {
		return nil, nil, fmt.Errorf("erro na conversão: %w", err)
	}

	return tx, converted, nil
}

func (s *TransactionService) GetLatestTransactions(limit int32) ([]*entities.Transaction, error) {
	ctx := context.Background()

	transactions, err := s.repo.GetLatestTransactions(ctx, limit)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar últimas transações: %w", err)
	}

	return transactions, nil
}
