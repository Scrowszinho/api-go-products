package database

import (
	"github.com/Scrowszinho/api-go-products/internal/dto"
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

func (b *Bet) FindBetsByUser(page, limit int, user_id int) ([]dto.GetSingleBetsByUser, error) {
	var (
		bets []dto.GetSingleBetsByUser
		err  error
	)
	if page != 0 && limit != 0 {
		err = b.DB.Table("bets").
			Select("bets.id as id, bets.amount as amount, outcomes.odds as odds, outcomes.status as status, outcomes.name as outcome_name, events.id as event_id ,events.name as event_name, events.start_time, events.end_time, events.description as event_description").
			Joins("JOIN outcomes ON bets.outcome_id = outcomes.id").
			Joins("JOIN events ON outcomes.event_id = events.id").
			Where("bets.user_id = ?", user_id).
			Limit(limit).
			Offset((page - 1) * limit).
			Scan(&bets).Error

	} else {
		err = b.DB.Table("bets").
			Select("bets.id as id, bets.amount as amount, outcomes.odds as odds, outcomes.status as status, outcomes.name as outcome_name, events.id as event_id ,events.name as event_name, events.start_time, events.end_time, events.description as event_description").
			Joins("JOIN outcomes ON bets.outcome_id = outcomes.id").
			Joins("JOIN events ON outcomes.event_id = events.id").
			Where("bets.user_id = ?", user_id).
			Scan(&bets).Error
	}
	return bets, err
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
