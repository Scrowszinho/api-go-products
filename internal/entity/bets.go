package entity

import "errors"

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

func NewBet(user *User, outcome_id int, amount float64, bonus float64, active bool) (*Bets, error) {
	if user.Balance < amount {
		return nil, errors.New("insufficient Balance")
	}
	bet := &Bets{
		UserID:    user.ID,
		OutcomeID: outcome_id,
		Amount:    amount,
		Active:    active,
		Bonus:     bonus,
	}
	return bet, nil
}
