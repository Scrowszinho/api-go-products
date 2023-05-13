package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMultipleBet(t *testing.T) {
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test", 1000)
	if err != nil {
		panic(err)
	}
	multiBet, err := CreateMultipleBet(user, 100)
	assert.Nil(t, err)
	assert.Equal(t, multiBet.Amount, 100.0)
}
