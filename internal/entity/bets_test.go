package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBet(t *testing.T) {
	bet, err := NewBet(1, 1, 100, 10.00, true)
	assert.Nil(t, err)
	calc := bet.Amount * 1.75
	assert.Equal(t, calc+(calc*(bet.Bonus/100)), 192.5)
}
