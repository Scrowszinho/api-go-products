package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCurrency(t *testing.T) {
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test", 1000)
	currency, err := CreateCurrency(100, DEPOSIT, time.Now(), user)
	assert.Nil(t, err)
	assert.Equal(t, currency.Amount+user.Balance, 1100.0)
	assert.NotNil(t, currency.Amount)
	currency, err = CreateCurrency(-1.0, DEPOSIT, time.Now(), user)
	assert.Error(t, err)
}
