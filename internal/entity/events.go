package entity

import (
	"errors"
	"time"
)

type Events struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	StartTime   time.Time `gorm:"not null" json:"start_time"`
	EndTime     time.Time `gorm:"not null" json:"end_time"`
	Description string    `json:"description"`
}

func NewEvent(name string, startTime time.Time, endTime time.Time, description string) (*Events, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if endTime.Before(startTime) {
		return nil, errors.New("final evente date muste be greater than or equal to the start date")
	}
	return &Events{
		Name:        name,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
	}, nil
}
