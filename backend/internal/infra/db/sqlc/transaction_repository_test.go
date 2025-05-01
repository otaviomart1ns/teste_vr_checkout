package sqlc_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/infra/db/sqlc"
)

func setupTestRepo(t *testing.T) *sqlc.TransactionRepository {
	t.Helper()

	cfg := config.Load()

	dbpool, err := pgxpool.New(context.Background(), cfg.PostgresURL)
	if err != nil {
		t.Fatalf("erro ao conectar no banco: %v", err)
	}

	return sqlc.NewTransactionRepository(dbpool)
}

func TestSaveAndFindByID(t *testing.T) {
	repo := setupTestRepo(t)

	tx := &entities.Transaction{
		ID:          uuid.NewString(),
		Description: "Teste Save",
		Date:        time.Now().UTC(),
		ValueUSD:    123.45,
	}

	err := repo.Save(context.Background(), tx)
	assert.NoError(t, err)

	saved, err := repo.FindByID(context.Background(), tx.ID)
	assert.NoError(t, err)
	assert.Equal(t, tx.Description, saved.Description)
	assert.Equal(t, tx.ValueUSD, saved.ValueUSD)
}

func TestGetLatestTransactions(t *testing.T) {
	repo := setupTestRepo(t)

	results, err := repo.GetLatestTransactions(context.Background(), 5)
	assert.NoError(t, err)
	assert.LessOrEqual(t, len(results), 5)
}

func TestSave_InvalidUUID(t *testing.T) {
	repo := setupTestRepo(t)

	tx := &entities.Transaction{
		ID:          "uuid-invalido",
		Description: "Teste erro UUID",
		Date:        time.Now(),
		ValueUSD:    50.0,
	}

	err := repo.Save(context.Background(), tx)
	assert.Error(t, err)
}

func TestFindByID_InvalidUUID(t *testing.T) {
	repo := setupTestRepo(t)

	tx, err := repo.FindByID(context.Background(), "uuid-invalido")
	assert.Error(t, err)
	assert.Nil(t, tx)
}

func TestFindByID_ForcedQueryError(t *testing.T) {
	repo := setupTestRepo(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tx, err := repo.FindByID(ctx, uuid.NewString())
	assert.Error(t, err)
	assert.Nil(t, tx)
}

func TestFindByID_NotFound(t *testing.T) {
	repo := setupTestRepo(t)

	nonExistentID := uuid.New()

	tx, err := repo.FindByID(context.Background(), nonExistentID.String())

	assert.NoError(t, err)
	assert.Nil(t, tx)
}

func TestGetLatestTransactions_DBError(t *testing.T) {
	repo := setupTestRepo(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	transactions, err := repo.GetLatestTransactions(ctx, 5)

	assert.Error(t, err)
	assert.Nil(t, transactions)
}
