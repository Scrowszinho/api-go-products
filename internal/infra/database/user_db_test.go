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
	user, _ := entity.NewUser("Test", "gustavo@gmail.com", "123456", "Test", 1000.0)
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
	userDB := NewUser(db)
	userFound, err := userDB.FindByEmail("1")
	assert.Nil(t, err)
	userFound1, err := userDB.FindByEmail("test@email.com")
	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, 1)
	assert.Equal(t, userFound.Email, "gustavo@gmail.com")
	assert.Equal(t, userFound.ID, userFound1.ID)
	assert.NotNil(t, userFound.Password)
}
