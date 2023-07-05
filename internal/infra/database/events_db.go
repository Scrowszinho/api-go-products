package database

import (
	"strconv"

	"github.com/Scrowszinho/api-go-products/internal/entity"
	"gorm.io/gorm"
)

type Events struct {
	DB *gorm.DB
}

func NewEvent(db *gorm.DB) *Events {
	return &Events{DB: db}
}

func (p *Events) Create(event *entity.Events) error {
	return p.DB.Create(event).Error
}

func (p *Events) FindById(id string) (*entity.Events, error) {
	var event entity.Events
	err := p.DB.First(&event, "id = ?", id).Error
	return &event, err
}

func (p *Events) FindByName(name string) (*entity.Events, error) {
	var event entity.Events
	err := p.DB.First(&event, "name = ?", name).Error
	return &event, err
}

func (e *Events) Update(event *entity.Events) error {
	id := strconv.Itoa(event.ID)
	_, err := e.FindById(id)
	if err != nil {
		return err
	}
	return e.DB.Save(event).Error
}

func (p *Events) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Events) FindAll(page, limit int, sort string) ([]entity.Events, error) {
	var (
		events []entity.Events
		err    error
	)

	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("start_time " + sort).Find(&events).Error
	} else {
		err = p.DB.Order("start_time" + sort).Find(&events).Error
	}
	return events, err
}
