package entity

import "errors"

type Bets struct {
	ID        int     `gorm:"primaryKey"`
	UserID    int     `gorm:"not null"`
	OutcomeID int     `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	Active    bool    `json:"active"`
	Bonus     float64
	Outcome   Outcome `gorm:"foreignKey:OutcomeID"`
}

func NewBet(user *User, outcome *Outcome, amount float64, bonus float64, active bool) (*Bets, error) {
	if user.Balance < amount {
		return nil, errors.New("insufficient Balance")
	}
	bet := &Bets{
		UserID:    user.ID,
		OutcomeID: outcome.ID,
		Amount:    amount,
		Active:    active,
		Bonus:     bonus,
	}
	return bet, nil
}
