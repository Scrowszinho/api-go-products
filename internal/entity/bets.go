package entity

import "gorm.io/gorm"

type Bets struct {
	ID                  int     `gorm:"primaryKey" json:"id"`
	UserID              int     `gorm:"not null" json:"user_id"`
	OutcomeID           int     `gorm:"not null" json:"outcome_id"`
	Amount              float64 `gorm:"not null" json:"amount"`
	Active              bool    `json:"active"`
	Payout              float64 `json:"payout"`
	Bonus               float64 `json:"bonus"`
	PayoutMultiplier    float64 `json:"payout_multiplier"`
	PotentialMultiplier float64 `json:"potential_multiplier"`
	CashoutMultiplier   float64 `json:"cashout_multiplier"`
	User                User    `gorm:"foreignKey:UserID"`
	Outcome             Outcome `gorm:"foreignKey:OutcomeID"`
	gorm.Model
}

func NewBet(userID int, outcome_id int, amount float64, bonus float64, active bool, odds float64) (*Bets, error) {
	bet := &Bets{
		UserID:           userID,
		OutcomeID:        outcome_id,
		Amount:           amount,
		Active:           active,
		Payout:           0,
		Bonus:            bonus,
		PayoutMultiplier: odds,
	}
	return bet, nil
}
