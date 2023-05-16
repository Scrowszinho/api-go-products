package database

import (
	"strconv"

	"github.com/Scrowszinho/api-go-products/internal/entity"
	"gorm.io/gorm"
)

type Outcome struct {
	DB *gorm.DB
}

func NewOutcome(db *gorm.DB) *Outcome {
	return &Outcome{DB: db}
}

func (o *Outcome) Create(outcome *entity.Outcome) error {
	return o.DB.Create(outcome).Error
}

func (o *Outcome) FindById(id string) (*entity.Outcome, error) {
	var outcome entity.Outcome
	if err := o.DB.Where("id = ? ", id).First(&outcome).Error; err != nil {
		return nil, err
	}
	return &outcome, nil
}

func (o *Outcome) Update(event *entity.Outcome) error {
	id := strconv.Itoa(event.ID)
	_, err := o.FindById(id)
	if err != nil {
		return err
	}
	return o.DB.Save(event).Error
}

func (o *Outcome) Delete(id string) error {
	outcome, err := o.FindById(id)
	if err != nil {
		return err
	}
	return o.DB.Delete(outcome).Error
}

func (o *Outcome) FindOutcomesByEventID(page, limit int, id string) ([]entity.Outcome, error) {
	var (
		outcomes []entity.Outcome
		err      error
	)

	if page != 0 && limit != 0 {
		err = o.DB.Limit(limit).Offset((page-1)*limit).Order("id "+"asc").Where("event_id = ?", id).Find(&outcomes).Error
	} else {
		err = o.DB.Order("id "+"asc").Where("event_id = ?", id).Error
	}
	return outcomes, err
}
