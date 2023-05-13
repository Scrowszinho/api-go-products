package database

import (
	"strconv"

	"github.com/Scrowszinho/api-go-products/internal/entity"
	"gorm.io/gorm"
)

type Event struct {
	DB *gorm.DB
}

func NewEvent(db *gorm.DB) *Event {
	return &Event{DB: db}
}

func (p *Event) Create(event *entity.Event) error {
	return p.DB.Create(event).Error
}

func (p *Event) FindById(id string) (*entity.Event, error) {
	var event entity.Event
	err := p.DB.First(&event, "id = ?", id).Error
	return &event, err
}

func (e *Event) Update(event *entity.Event) error {
	id := strconv.Itoa(event.ID)
	_, err := e.FindById(id)
	if err != nil {
		return err
	}
	return e.DB.Save(event).Error
}

func (p *Event) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Event) FindAll(page, limit int, sort string) ([]entity.Event, error) {
	var (
		events []entity.Event
		err    error
	)

	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&events).Error
	} else {
		err = p.DB.Order("created_at" + sort).Find(&events).Error
	}
	return events, err
}
