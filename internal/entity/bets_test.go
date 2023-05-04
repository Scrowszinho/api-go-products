package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateBet(t *testing.T) {
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test", 1000)
	event, err := NewEvent("Test", time.Now(), time.Now().AddDate(0, 0, 2), "Test")
	assert.Nil(t, err)
	outcome, err := CreateOutcome(event, "Test", 1.75)
	bet, err := NewBet(user, outcome, 100, CASHOUT, 10.00)
	calc := bet.Amount * outcome.Odds
	assert.Equal(t, calc+(calc*(bet.Bonus/100)), 192.5)
	assert.Nil(t, err)
	bet, err = NewBet(user, outcome, 2000, CASHOUT, 15.00)
	assert.Error(t, err)
}
