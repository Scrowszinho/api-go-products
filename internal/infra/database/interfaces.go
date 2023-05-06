package database

import "github.com/Scrowszinho/api-go-products/internal/entity"

type User interface {
	CreateUser(user *entity.User) error
	FindByEmailOrUsername(name string) (*User, error)
}

type Bets interface {
	Create(product *entity.Bets) error
	FindAll(page, limit int, sort string) ([]entity.Bets, error)
	FindById(id uint) (*entity.Bets, error)
	Update(bet *entity.Bets) error
	Delete(id string) error
}
