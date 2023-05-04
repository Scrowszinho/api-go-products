package entity

import (
	"errors"
	"time"
)

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Description string
}

func NewEvent(name string, startTime time.Time, endTime time.Time, description string) (*Event, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}
	if endTime.Before(startTime) {
		return nil, errors.New("Final evente date muste be greater than or equal to the start date")
	}
	return &Event{
		Name:        name,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
	}, nil
}
