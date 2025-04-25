package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/otaviomart1ns/teste_vr_checkout/internal/domain"
	"github.com/otaviomart1ns/teste_vr_checkout/internal/infra/db/sqlc"
)

type TransactionRepository struct {
	q *sqlc.Queries
}

func NewTransactionRepository(q *sqlc.Queries) *TransactionRepository {
	return &TransactionRepository{q: q}
}

func (r *TransactionRepository) Save(t *domain.Transaction) error {
	params := sqlc.CreateTransactionParams{
		ID:          uuid.MustParse(t.ID),
		Description: t.Description,
		Date:        t.Date,
		Amount:      t.Amount,
	}
	return r.q.CreateTransaction(context.Background(), params)
}

func (r *TransactionRepository) GetByID(id string) (*domain.Transaction, error) {
	txID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	sqlcTx, err := r.q.GetTransaction(context.Background(), txID)
	if err != nil {
		return nil, err
	}

	return &domain.Transaction{
		ID:          sqlcTx.ID.String(),
		Description: sqlcTx.Description,
		Date:        sqlcTx.Date,
		Amount:      sqlcTx.Amount,
	}, nil
}
