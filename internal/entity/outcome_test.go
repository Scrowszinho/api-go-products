package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateOutcome(t *testing.T) {
	event, err := NewEvent("Test", time.Now(), time.Now(), "Test")
	outcomes, err := CreateOutcome(event, "Test", 2.75)
	assert.Nil(t, err)
	assert.NotNil(t, outcomes)
	assert.Equal(t, "Test", outcomes.Name)
}

func TestCreateMultiOutcome(t *testing.T) {
	event, err := NewEvent("Test", time.Now(), time.Now(), "Test")
	outcome, err := CreateOutcome(event, "Test", 2.75)
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test", 1000)
	multBet, err := CreateMultipleBet(user, 100)
	outcomes, err := CreateMultiOutcome(multBet, outcome)
	assert.Nil(t, err)
	assert.Equal(t, outcomes.Odds*multBet.Amount, 275.0)

}
