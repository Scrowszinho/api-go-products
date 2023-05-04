package entity

import "errors"

type Bets struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	OutcomeID uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	Status    string
	Bonus     float64
	Outcome   Outcome `gorm:"foreignKey:OutcomeID"`
}

func NewBet(user *User, outcome *Outcome, amount float64, status BetStatus, bonus float64) (*Bets, error) {
	if user.Balance < amount {
		return nil, errors.New("Insufficient Balance")
	}
	bet := &Bets{
		UserID:    user.ID,
		Status:    string(status),
		OutcomeID: outcome.ID,
		Amount:    amount,
		Bonus:     bonus,
	}
	return bet, nil
}
