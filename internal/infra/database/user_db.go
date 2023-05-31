package database

import (
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmailOrNickname(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ? OR nickname = ?", email, email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindById(id int) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
