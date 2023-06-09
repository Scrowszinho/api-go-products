package database

import (
	"github.com/Scrowszinho/api-go-products/internal/dto"
	"github.com/Scrowszinho/api-go-products/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type EventsInterface interface {
	Create(product *entity.Events) error
	FindAll(page, limit int, sort string) ([]entity.Events, error)
	FindById(id string) (*entity.Events, error)
	Update(event *entity.Events) error
	Delete(id string) error
}

type BetsInterface interface {
	Create(bets *entity.Bets) error
	FindById(id int) (*entity.Bets, error)
	FindBetsByUser(page, limit int, user_id int) ([]dto.GetSingleBetsByUser, error)
	Update(bets *entity.Bets) error
	Delete(id int) error
}

type OutcomeInterface interface {
	Create(outcome *entity.Outcome) error
	FindById(id int) (*entity.Outcome, error)
	FindOutcomesById(id int, page, limit int, sort string) ([]entity.Outcome, error)
	Update(bets *entity.Outcome) error
	Delete(id int) error
}

type MultipleOutcomesInterface interface {
}
