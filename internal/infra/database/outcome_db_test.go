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
	eventDB := NewEvent(db)
	outcomeDB := NewOutcome(db)
	event, err := eventDB.FindById("1")
	if err != nil {
		panic(err)
	}
	outcome, err := entity.CreateOutcome(event, "Teste", 1.75, entity.AVOIDED)
	err = outcomeDB.Create(outcome)
	assert.Nil(t, err)
	assert.Equal(t, outcome.EventID, event.ID)
}

func TestFindOutcomeById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	outcomeDB := NewOutcome(db)

	outcome, err := outcomeDB.FindById("1")
	assert.Nil(t, err)
	assert.Equal(t, outcome.Name, "Teste")
	assert.Equal(t, outcome.Odds, 1.75)
}

func TestDeleteOutcomeById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	outcomeDB := NewOutcome(db)

	err := outcomeDB.Delete("3")
	assert.Nil(t, err)
	err = outcomeDB.Delete("0")
	assert.Error(t, err)
}

func TestUpdateOutcome(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	eventDB := NewEvent(db)
	outcomeDB := NewOutcome(db)
	event, err := eventDB.FindById("4")
	if err != nil {
		panic(err)
	}

	outcome := entity.Outcome{ID: 4, Name: "Milan vencer", Odds: 2, EventID: event.ID, Event: *event}
	err = outcomeDB.Update(&outcome)
	assert.Nil(t, err)
}
