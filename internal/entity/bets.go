package entity

type Bets struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	UserID    int     `gorm:"not null" json:"user_id"`
	OutcomeID int     `gorm:"not null" json:"outcome_id"`
	Amount    float64 `gorm:"not null" json:"amount"`
	Active    bool    `json:"active"`
	Bonus     float64 `json:"bonus"`
	User      User    `gorm:"foreignKey:UserID"`
	Outcome   Outcome `gorm:"foreignKey:OutcomeID"`
}

func NewBet(userID int, outcome_id int, amount float64, bonus float64, active bool) (*Bets, error) {
	bet := &Bets{
		UserID:    userID,
		OutcomeID: outcome_id,
		Amount:    amount,
		Active:    active,
		Bonus:     bonus,
	}
	return bet, nil
}
