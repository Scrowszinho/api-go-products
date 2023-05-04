package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	ID          uint      `gorm:"primaryKey"`
	Amount      float64   `json:"amount"`
	DateCreated time.Time `json:"date_created"`
	UserID      uint      `gorm:"not null"`
	Status      CurrencyStatus
	User        User
	gorm.Model
}

func CreateCurrency(amount float64, status CurrencyStatus, dateCreated time.Time, user *User) (*Currency, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	return &Currency{
		Amount:      amount,
		Status:      status,
		DateCreated: dateCreated,
		UserID:      user.ID,
	}, nil
}
