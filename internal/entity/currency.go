package entity

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	Value       float64   `json:"value"`
	Type        string    `json:"type"`
	GamblingID  int       `json:"gambling_id"`
	DateCreated time.Time `json:"date_created"`
	UserID      int       `json:"user_id"`
	gorm.Model
}

func CreateCurrency(value float64, typee string, user_id int, dateCreated time.Time, gambling_id int) *Currency {
	return &Currency{
		Value:       value,
		Type:        typee,
		UserID:      user_id,
		GamblingID:  gambling_id,
		DateCreated: dateCreated,
	}
}
