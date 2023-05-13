package entity

type Outcome struct {
	ID      int     `gorm:"primaryKey"`
	Name    string  `gorm:"not null"`
	Odds    float64 `gorm:"not null"`
	EventID int     `gorm:"not null"`
	Event   Event   `gorm:"foreignKey:EventID"`
}

type MultiOutcome struct {
	ID         int      `gorm:"primaryKey"`
	MultiBetID int      `gorm:"not null"`
	OutcomeID  int      `gorm:"not null"`
	Odds       float64  `gorm:"not null"`
	MultiBet   MultiBet `gorm:"foreignKey:MultiBetID"`
	Outcome    Outcome  `gorm:"foreignKey:OutcomeID"`
}

func CreateOutcome(event *Event, name string, odds float64) (*Outcome, error) {
	return &Outcome{
		EventID: event.ID,
		Name:    name,
		Odds:    odds,
	}, nil
}

func CreateMultiOutcome(multiBet *MultiBet, outcome *Outcome) (*MultiOutcome, error) {
	return &MultiOutcome{
		MultiBetID: multiBet.ID,
		OutcomeID:  outcome.ID,
		Odds:       outcome.Odds,
	}, nil
}
