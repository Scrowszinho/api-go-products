package entity

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Description string
}

func NewEvent(name string, startTime time.Time, endTime time.Time, description string) (*Event, error) {
	return &Event{
		Name:        name,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
	}, nil
}
