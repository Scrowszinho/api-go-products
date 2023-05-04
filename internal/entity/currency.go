package entity

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	ID          uint      `gorm:"primaryKey"`
	Amount      float64   `json:"amount"`
	Type        string    `gorm:"not null"`
	DateCreated time.Time `json:"date_created"`
	UserID      uint      `gorm:"not null"`
	User        User
	gorm.Model
}

func CreateCurrency(amount float64, typee string, dateCreated time.Time, user *User) *Currency {
	return &Currency{
		Amount:      amount,
		Type:        typee,
		UserID:      user.ID,
		DateCreated: dateCreated,
	}
}
