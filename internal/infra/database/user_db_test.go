package database

import (
	"fmt"
	"testing"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	config, err := configs.GetConfigs(".")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Gustavo", "Test", "gustavo@gmail.com", "123456", "Test", 0.0)
	userDB := NewUser(db)

	err = userDB.Create(user)
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
	config, err := configs.GetConfigs(".")
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Gustavo", "Test", "gustavo@gmail.com", "123456", "Test", 1000.0)
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}
