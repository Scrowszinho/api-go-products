package entity

import "errors"

type Outcome struct {
	ID      int       `gorm:"primaryKey" json:"id"`
	Name    string    `gorm:"not null" json:"name"`
	Odds    float64   `gorm:"not null" json:"odds"`
	EventID int       `gorm:"not null" json:"event_id"`
	Status  BetStatus `gorm:"not null" json:"status"`
	Event   Event     `gorm:"foreignKey:EventID"`
}

type MultiOutcome struct {
	ID          int       `gorm:"primaryKey"`
	MultiBetsID int       `gorm:"not null"`
	OutcomeID   int       `gorm:"not null"`
	Odds        float64   `gorm:"not null"`
	Status      BetStatus `gorm:"not null"`
	MultiBets   MultiBets `gorm:"foreignKey:MultiBetsID"`
	Outcome     Outcome   `gorm:"foreignKey:OutcomeID"`
}

func CreateOutcome(event int, name string, odds float64, status BetStatus) (*Outcome, error) {
	if event == 0 {
		return nil, errors.New("event not selected")
	}
	return &Outcome{
		EventID: event,
		Name:    name,
		Odds:    odds,
		Status:  status,
	}, nil
}

func CreateMultiOutcome(multiBet *MultiBets, outcome *Outcome) (*MultiOutcome, error) {
	return &MultiOutcome{
		MultiBetsID: multiBet.ID,
		OutcomeID:   outcome.ID,
		Odds:        outcome.Odds,
		Status:      outcome.Status,
	}, nil
}
