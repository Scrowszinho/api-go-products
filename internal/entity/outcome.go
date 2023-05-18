package entity

type Outcome struct {
	ID      int       `gorm:"primaryKey"`
	Name    string    `gorm:"not null"`
	Odds    float64   `gorm:"not null"`
	EventID int       `gorm:"not null"`
	Status  BetStatus `gorm:"not null"`
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

func CreateOutcome(event *Event, name string, odds float64, status BetStatus) (*Outcome, error) {
	return &Outcome{
		EventID: event.ID,
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
