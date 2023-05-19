package entity

import "errors"

type MultiBets struct {
	ID     int     `gorm:"primaryKey"`
	UserID int     `gorm:"not null"`
	Amount float64 `gorm:"not null"`
	User   User    `gorm:"foreignKey:UserID"`
}

func CreateMultipleBets(user *User, amount float64) (*MultiBets, error) {
	if user.Balance < amount {
		return nil, errors.New("insufficient Balance")
	}
	return &MultiBets{
		UserID: user.ID,
		Amount: amount,
	}, nil
}
