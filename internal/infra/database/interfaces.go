package database

import "github.com/Scrowszinho/api-go-products/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmailOrNickname(email string) (*User, error)
}

type Events interface {
	Create(product *entity.Event) error
	FindAll(page, limit int, sort string) ([]entity.Event, error)
	FindById(id int) (*entity.Event, error)
	Update(event *entity.Event) error
	Delete(id string) error
}

type Bets interface {
	Create(bets *entity.Bets) error
	FindAll(page, limit int, sort string) ([]entity.Bets, error)
	FindById(id int) (*entity.Bets, error)
	Update(bets *entity.Bets) error
	Delete(id string) error
}

type Outcome interface {
}
