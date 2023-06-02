package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOutcome(t *testing.T) {
	outcomes, err := CreateOutcome(1, "Test", 2.75, AVOIDED)
	assert.Nil(t, err)
	assert.NotNil(t, outcomes)
	assert.Equal(t, "Test", outcomes.Name)
}

func TestCreateMultiBetsOutcome(t *testing.T) {
	outcome, err := CreateOutcome(1, "Test", 2.75, CASHOUT)
	if err != nil {
		panic(err)
	}
	multBet, err := CreateMultipleBets(1, 100)
	if err != nil {
		panic(err)
	}
	outcomes, err := CreateMultiBetsOutcome(multBet, outcome)
	assert.Nil(t, err)
	assert.Equal(t, outcomes.Odds*multBet.Amount, 275.0)

}
