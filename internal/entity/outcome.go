package entity

type Outcome struct {
	ID       int     `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"not null" json:"name"`
	Odds     float64 `gorm:"not null" json:"odds"`
	MarketID int     `gorm:"not null" json:"market_id"`
	Market   Market  `gorm:"foreignKey:MarketID"`
}

func CreateOutcome(name string, odds float64, status BetStatus, marketId int) (*Outcome, error) {
	return &Outcome{
		Name:     name,
		Odds:     odds,
		MarketID: marketId,
	}, nil
}
