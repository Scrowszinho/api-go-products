package database

import "github.com/Scrowszinho/api-go-products/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*User, error)
}

type Events interface {
	Create(product *entity.Event) error
	FindAll(page, limit int, sort string) ([]entity.Event, error)
	FindById(id int) (*entity.Event, error)
	Update(event *entity.Event) error
	Delete(id string) error
}
