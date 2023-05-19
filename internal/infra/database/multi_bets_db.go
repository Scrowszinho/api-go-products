package database

import (
	"strconv"

	"github.com/Scrowszinho/api-go-products/internal/entity"
	"gorm.io/gorm"
)

type MultiBet struct {
	DB *gorm.DB
}

func NewMultiBet(db *gorm.DB) *MultiBet {
	return &MultiBet{DB: db}
}

func (b *MultiBet) Create(bets *entity.MultiBets) error {
	return b.DB.Create(bets).Error
}

func (b *MultiBet) FindById(id string) (*entity.MultiBets, error) {
	var bets entity.MultiBets
	err := b.DB.First(&bets, "id = ?", id).Error
	return &bets, err
}

func (b *MultiBet) Update(bets *entity.MultiBets) error {
	id := strconv.Itoa(bets.ID)
	_, err := b.FindById(id)
	if err != nil {
		return err
	}
	return b.DB.Save(bets).Error
}

func (b *MultiBet) Delete(id string) error {
	bets, err := b.FindById(id)
	if err != nil {
		return err
	}
	return b.DB.Delete(bets).Error
}
