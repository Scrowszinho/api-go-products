package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMultipleBets(t *testing.T) {
	multiBet, err := CreateMultipleBets(1, 100)
	assert.Nil(t, err)
	assert.Equal(t, multiBet.Amount, 100.0)
}
