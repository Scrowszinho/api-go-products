package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBet(t *testing.T) {
	user, err := NewUser("Test", "Test", "test@gmail", "123", "Test", 100)
	event, err := NewEvent("Test", time.Now(), time.Now(), "Test")
	outcomes, err := CreateOutcome(event, "Test", 2.75)
	bet, err := NewBet(user, outcomes, 100.5, SOLVED, 0)
	assert.Nil(t, err)
	assert.NotNil(t, bet)
	assert.NotEmpty(t, bet.OutcomeID)
}
