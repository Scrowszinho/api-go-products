package entity

import "gorm.io/gorm"

type Bets struct {
	ID                  int       `gorm:"primaryKey" json:"id"`
	Amount              float64   `gorm:"not null" json:"amount"`
	Active              bool      `json:"active"`
	Payout              float64   `json:"payout"`
	Bonus               float64   `json:"bonus"`
	Status              BetStatus `json:"status"`
	PayoutMultiplier    float64   `json:"payout_multiplier"`
	PotentialMultiplier float64   `json:"potential_multiplier"`
	CashoutMultiplier   float64   `json:"cashout_multiplier"`
	UserID              int       `gorm:"not null" json:"user_id"`
	EventID             int       `gorm:"not null" json:"event_id"`
	User                User      `gorm:"foreignKey:UserID"`
	Outcome             []Outcome `gorm:"foreignKey:OutcomeID"`
	gorm.Model
}

func NewBet(userID int, eventsID int, amount float64, bonus float64, active bool) (*Bets, error) {
	bet := &Bets{
		UserID:  userID,
		EventID: eventsID,
		Amount:  amount,
		Active:  active,
		Payout:  0,
		Bonus:   bonus,
	}
	return bet, nil
}
