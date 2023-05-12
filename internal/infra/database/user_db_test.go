package database

import (
	"testing"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/Scrowszinho/api-go-products/migrations"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	user, _ := entity.NewUser("Gustavo", "Test", "gustavo@gmail.com", "123456", "Test", 1000.0)
	userDB := NewUser(db)

	err := userDB.Create(user)
	assert.Nil(t, err)

	var userFounder entity.User
	err = db.First(&userFounder, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, userFounder.ID, user.ID)
	assert.Equal(t, userFounder.Email, user.Email)
	assert.Equal(t, userFounder.Name, user.Name)
	assert.NotNil(t, userFounder.Password)
}

func TestByEmail(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	user, _ := entity.NewUser("Gustavo", "Test", "gustavo@gmail.com", "123456", "Test", 1000.0)
	userDB := NewUser(db)

	err := userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}
