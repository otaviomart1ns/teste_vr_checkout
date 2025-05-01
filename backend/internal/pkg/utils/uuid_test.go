package utils_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	id := utils.GenerateUUID()
	_, err := uuid.Parse(id)
	assert.NoError(t, err)
}
