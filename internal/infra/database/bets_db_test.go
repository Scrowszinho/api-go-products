package database

import (
	"testing"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/Scrowszinho/api-go-products/migrations"
	"github.com/stretchr/testify/assert"
)

func TestCreateBet(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betsDB := NewBet(db)
	bets, err := entity.NewBet(1, 1, 100, 0, true)
	err = betsDB.Create(bets)
	assert.Nil(t, err)
}

func TestFindBetById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betDB := NewBet(db)

	bet, err := betDB.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, bet.Amount, 100.0)
}

func TestUpdateBet(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betsDB := NewBet(db)
	userDB := NewUser(db)
	outcomeDB := NewOutcome(db)
	user, err := userDB.FindByEmail("gustavo@gmail.com")
	if err != nil {
		panic(err)
	}
	outcome, err := outcomeDB.FindById(1)
	if err != nil {
		panic(err)
	}
	bets := entity.Bets{ID: 1, UserID: 1, OutcomeID: 1, Amount: 10.0, User: *user, Active: true, Bonus: 0, Outcome: *outcome}

	err = betsDB.Update(&bets)
	assert.Nil(t, err)
}

func TestDeleteBetById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betDB := NewBet(db)

	err := betDB.Delete(0)
	assert.Error(t, err)
}
