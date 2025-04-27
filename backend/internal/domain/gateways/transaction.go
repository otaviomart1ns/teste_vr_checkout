package gateways

import (
	"context"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
)

type TransactionProducer interface {
	PublishTransaction(ctx context.Context, tx *entities.Transaction) error
}

type TransactionRepository interface {
	Save(ctx context.Context, tx *entities.Transaction) error
}
