package utils_test

import (
	"testing"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsAlphanumeric(t *testing.T) {
	assert.True(t, utils.IsAlphanumeric("abc123"))
	assert.True(t, utils.IsAlphanumeric("ABC 123 xyz"))
	assert.False(t, utils.IsAlphanumeric("abc@123"))
	assert.False(t, utils.IsAlphanumeric("with\nnewline"))
	assert.False(t, utils.IsAlphanumeric("tab\tseparated"))
}

func TestRoundToCents(t *testing.T) {
	assert.Equal(t, 12.34, utils.RoundToCents(12.337))
	assert.Equal(t, 12.35, utils.RoundToCents(12.345))
	assert.Equal(t, 0.0, utils.RoundToCents(0.004))
	assert.Equal(t, 0.01, utils.RoundToCents(0.005))
}
