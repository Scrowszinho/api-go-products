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
	bet, err := NewBet(user, 1, 100, 10.00, true)
	assert.Nil(t, err)
	calc := bet.Amount * 1.75
	assert.Equal(t, calc+(calc*(bet.Bonus/100)), 192.5)
}
