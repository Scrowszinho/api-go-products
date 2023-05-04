package entity

import "errors"

type MultiBet struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"not null"`
	Amount float64 `gorm:"not null"`
	User   User    `gorm:"foreignKey:UserID"`
}

func CreateMultipleBet(user *User, amount float64) (*MultiBet, error) {
	if user.Balance < amount {
		return nil, errors.New("Insufficient Balance")
	}
	return &MultiBet{
		UserID: user.ID,
		Amount: amount,
	}, nil
}
