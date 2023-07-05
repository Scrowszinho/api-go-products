package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	event, err := NewEvent("Test", time.Now(), time.Now().Add(time.Hour*24), "Test")
	assert.Nil(t, err)
	assert.NotNil(t, event)
}
