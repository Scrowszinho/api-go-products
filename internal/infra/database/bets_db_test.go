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
	userDB := NewUser(db)
	outcomeDB := NewOutcome(db)
	user, err := userDB.FindByEmailOrNickname("gustavo@gmail.com")
	if err != nil {
		panic(err)
	}
	outcome, err := outcomeDB.FindById("4")
	if err != nil {
		panic(err)
	}
	bets, err := entity.NewBet(user, outcome, 100, entity.AVOIDED, 0, true)
	err = betsDB.Create(bets)
	assert.Nil(t, err)
}

func TestFindBetById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betDB := NewBet(db)

	bet, err := betDB.FindById("1")
	assert.Nil(t, err)
	assert.Equal(t, bet.Amount, 100.0)
	assert.Equal(t, bet.Status, "AVOIDED")
}

func TestDeleteBetById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betDB := NewBet(db)

	err := betDB.Delete("2")
	assert.Nil(t, err)
}
