package database

import (
	"testing"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/Scrowszinho/api-go-products/migrations"
	"github.com/stretchr/testify/assert"
)

func TestCreateOutcome(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	outcomeDB := NewOutcome(db)
	outcome, err := entity.CreateOutcome("Lutado A vencera", 1.75, entity.AVOIDED, 1)
	err = outcomeDB.Create(outcome)
	assert.Nil(t, err)
	assert.Equal(t, outcome.Name, "Lutado A vencera")
}

func TestFindOutcomeById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	outcomeDB := NewOutcome(db)

	outcome, err := outcomeDB.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, outcome.Name, "Lutado A vencera")
	assert.Equal(t, outcome.Odds, 1.75)
}

func TestUpdateOutcome(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	outcomeDB := NewOutcome(db)

	outcome := entity.Outcome{ID: 1, Name: "Milan vencer", Odds: 2, MarketID: 1}
	err := outcomeDB.Update(&outcome)
	assert.Nil(t, err)
}

func TestDeleteOutcomeById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	outcomeDB := NewOutcome(db)

	err := outcomeDB.Delete(0)
	assert.Error(t, err)
}
