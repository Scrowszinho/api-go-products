package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBet(t *testing.T) {
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test", 1000)
	if err != nil {
		panic(err)
	}
	outcome, err := CreateOutcome(1, "Test", 1.75, AVOIDED)
	if err != nil {
		panic(err)
	}
	bet, err := NewBet(user, outcome, 100, 10.00, true)
	assert.Nil(t, err)
	calc := bet.Amount * outcome.Odds
	assert.Equal(t, calc+(calc*(bet.Bonus/100)), 192.5)
}
