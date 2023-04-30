package entity

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	Value       float64   `json:"value"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
	gorm.Model
}

func CreateCurrency(value float64, typee string, status string, dateCreated time.Time) *Currency {
	return &Currency{
		Value:       value,
		Type:        typee,
		Status:      status,
		DateCreated: dateCreated,
	}
}
