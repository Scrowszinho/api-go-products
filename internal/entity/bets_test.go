package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBet(t *testing.T) {
	bet, err := NewSimpleBet("Test", "Test", "Test", "SOLVED", 1.75, 1.75, 10)
	assert.Nil(t, err)
	assert.NotNil(t, bet)
	assert.NotEmpty(t, bet.Odd)
	assert.NotEmpty(t, bet.Value)
	assert.NotEmpty(t, bet.Name)
	assert.NotEmpty(t, bet.Status)
	assert.Equal(t, bet.Value*bet.Odd, 17.5)
}
