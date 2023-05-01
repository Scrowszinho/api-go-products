package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Nickname)
	assert.Equal(t, "Test", user.Name)
	assert.Equal(t, "test@gmail.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Test", "Test", "test@gmail.com", "123456", "Test")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}
