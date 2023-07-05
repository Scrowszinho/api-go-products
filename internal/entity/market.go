package entity

import (
	"errors"

	"gorm.io/gorm"
)

type Market struct {
	Name string `gorm:"not null" json:"name"`
	gorm.Model
}

func CreateMarket(name string) (*Market, error) {
	if name == "" {
		return nil, errors.New("provide.name")
	}
	return &Market{
		Name: name,
	}, nil
}
