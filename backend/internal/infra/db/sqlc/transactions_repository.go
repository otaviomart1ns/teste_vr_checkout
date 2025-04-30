package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
)

type TransactionRepository struct {
	q *Queries
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		q: New(db),
	}
}

func (r *TransactionRepository) Save(ctx context.Context, tx *entities.Transaction) error {
	uid, err := uuid.Parse(tx.ID)
	if err != nil {
		return err
	}

	_, err = r.q.CreateTransaction(ctx, CreateTransactionParams{
		ID:          uid,
		Description: tx.Description,
		Date:        tx.Date,
		Amount:      tx.ValueUSD,
	})

	return err
}

func (r *TransactionRepository) FindByID(ctx context.Context, id string) (*entities.Transaction, error) {
	txID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	res, err := r.q.GetTransaction(ctx, txID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &entities.Transaction{
		ID:          res.ID.String(),
		Description: res.Description,
		Date:        res.Date,
		ValueUSD:    res.Amount,
	}, nil
}

func (r *TransactionRepository) GetLatestTransactions(ctx context.Context, limit int32) ([]*entities.Transaction, error) {
	results, err := r.q.GetLatestTransactions(ctx, limit)
	if err != nil {
		return nil, err
	}

	var transactions []*entities.Transaction
	for _, res := range results {
		transactions = append(transactions, &entities.Transaction{
			ID:          res.ID.String(),
			Description: res.Description,
			Date:        res.Date,
			ValueUSD:    res.Amount,
		})
	}

	return transactions, nil
}