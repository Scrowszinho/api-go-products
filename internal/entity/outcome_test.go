package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateOutcome(t *testing.T) {
	event, err := NewEvent("Test", time.Now(), time.Now(), "Test")
	if err != nil {
		panic(err)
	}
	outcomes, err := CreateOutcome(event, "Test", 2.75, AVOIDED)
	assert.Nil(t, err)
	assert.NotNil(t, outcomes)
	assert.Equal(t, "Test", outcomes.Name)
}

func TestCreateMultiOutcome(t *testing.T) {
	event, err := NewEvent("Test", time.Now(), time.Now(), "Test")
	if err != nil {
		panic(err)
	}
	outcome, err := CreateOutcome(event, "Test", 2.75, CASHOUT)
	if err != nil {
		panic(err)
	}
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test", 1000)
	if err != nil {
		panic(err)
	}
	multBet, err := CreateMultipleBets(user, 100)
	if err != nil {
		panic(err)
	}
	outcomes, err := CreateMultiOutcome(multBet, outcome)
	assert.Nil(t, err)
	assert.Equal(t, outcomes.Odds*multBet.Amount, 275.0)

}
