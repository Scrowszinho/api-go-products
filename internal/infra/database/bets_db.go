package database

import (
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"gorm.io/gorm"
)

type Bet struct {
	DB *gorm.DB
}

func NewBet(db *gorm.DB) *Bet {
	return &Bet{DB: db}
}

func (b *Bet) Create(bets *entity.Bets) error {
	return b.DB.Create(bets).Error
}

func (b *Bet) FindById(id int) (*entity.Bets, error) {
	var bets entity.Bets
	err := b.DB.First(&bets, "id = ?", id).Error
	return &bets, err
}

func (b *Bet) Update(bets *entity.Bets) error {
	_, err := b.FindById(bets.ID)
	if err != nil {
		return err
	}
	return b.DB.Save(bets).Error
}

func (b *Bet) Delete(id int) error {
	bets, err := b.FindById(id)
	if err != nil {
		return err
	}
	return b.DB.Delete(bets).Error
}

func (b *Bet) FindAll(page, limit int, sort string) ([]entity.Bets, error) {
	var (
		bets []entity.Bets
		err  error
	)

	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = b.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&bets).Error
	} else {
		err = b.DB.Order("created_at" + sort).Find(&bets).Error
	}
	return bets, err
}
