package database

import (
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

func (o *Outcome) FindById(id int) (*entity.Outcome, error) {
	var outcome entity.Outcome
	if err := o.DB.Where("id = ? ", id).First(&outcome).Error; err != nil {
		return nil, err
	}
	return &outcome, nil
}

func (o *Outcome) FindOutcomesById(id int, page, limit int, sort string) ([]entity.Outcome, error) {
	var (
		outcomes []entity.Outcome
		err      error
	)

	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = o.DB.Limit(limit).Offset((page-1)*limit).Where("event_id = ?", id).Order("id " + sort).Find(&outcomes).Error
	} else {
		err = o.DB.Where("event_id = ?", id).Order("id" + sort).Find(&outcomes).Error
	}
	return outcomes, err
}

func (o *Outcome) Update(outcome *entity.Outcome) error {

	_, err := o.FindById(outcome.ID)
	if err != nil {
		return err
	}
	return o.DB.Save(outcome).Error
}

func (o *Outcome) Delete(id int) error {
	outcome, err := o.FindById(id)
	if err != nil {
		return err
	}
	return o.DB.Delete(outcome).Error
}
