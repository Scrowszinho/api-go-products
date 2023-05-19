package database

import (
	"testing"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/Scrowszinho/api-go-products/migrations"
	"github.com/stretchr/testify/assert"
)

func TestCreateMultiBet(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betsDB := NewMultiBet(db)
	userDB := NewUser(db)
	user, err := userDB.FindByEmailOrNickname("gustavo@gmail.com")
	if err != nil {
		panic(err)
	}
	bets, err := entity.CreateMultipleBets(user, 100.0)
	if err != nil {
		panic(err)
	}
	err = betsDB.Create(bets)
	assert.Nil(t, err)
}

func TestFindBetMultiById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betDB := NewBet(db)

	bet, err := betDB.FindById("1")
	assert.Nil(t, err)
	assert.Equal(t, bet.Amount, 100.0)
	assert.Equal(t, bet.Status, "AVOIDED")
}

func TestDeleteMultiBetById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	betDB := NewBet(db)

	err := betDB.Delete("2")
	assert.Nil(t, err)
}
