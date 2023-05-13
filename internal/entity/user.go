package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Nickname string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Balance  float64
	gorm.Model
}

func NewUser(name string, lastName string, email string, password string, nickname string, balance float64) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     name,
		LastName: lastName,
		Email:    email,
		Nickname: nickname,
		Password: string(hash),
		Balance:  balance,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
