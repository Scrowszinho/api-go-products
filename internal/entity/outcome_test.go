package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOutcome(t *testing.T) {
	outcomes, err := CreateOutcome("Lutador A", 1.75, AVOIDED, 1)
	assert.Nil(t, err)
	assert.NotNil(t, outcomes)
	assert.Equal(t, "Lutador A", outcomes.Name)
}
