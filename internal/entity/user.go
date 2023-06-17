package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `gorm:"primaryKey"  json:"id"`
	Name      string    `gorm:"not null"  json:"name"`
	Email     string    `gorm:"unique;not null"  json:"email"`
	Nickname  string    `gorm:"unique;not null"  json:"nickname"`
	Password  string    `gorm:"not null"  json:"password"`
	Balance   float64   ` json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name string, email string, password string, nickname string, balance float64) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     name,
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
