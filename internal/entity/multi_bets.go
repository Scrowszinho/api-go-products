package entity

import (
	"gorm.io/gorm"
)

type MultiBets struct {
	ID     int       `gorm:"primaryKey" json:"id"`
	UserID int       `gorm:"not null" json:"user_id"`
	Amount float64   `gorm:"not null" json:"amount"`
	Status BetStatus `gorm:"not null" json:"status"`
	gorm.Model
}

func CreateMultipleBets(user_id int, amount float64) (*MultiBets, error) {

	return &MultiBets{
		UserID: user_id,
		Amount: amount,
	}, nil
}
