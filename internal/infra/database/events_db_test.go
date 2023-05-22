package database

import (
	"testing"
	"time"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/Scrowszinho/api-go-products/migrations"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	event, _ := entity.NewEvent("Ufc", time.Now(), time.Now().AddDate(0, 0, 1), "Test")
	eventDB := NewEvent(db)

	err := eventDB.Create(event)
	assert.Nil(t, err)
	assert.NotEmpty(t, event.ID)
	assert.NotEqual(t, event.StartTime, event.EndTime)
}

func TestFidById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	eventDB := NewEvent(db)

	event, err := eventDB.FindById("1")
	assert.Nil(t, err)
	assert.Equal(t, event.Name, "Ufc")
	assert.Equal(t, event.Description, "Test")
}

func TestFidByName(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	eventDB := NewEvent(db)

	event, err := eventDB.FindByName("Ufc")
	assert.Nil(t, err)
	assert.Equal(t, event.Name, "Ufc")
	assert.Equal(t, event.Description, "Test")
}

func TestUpdate(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	eventDB := NewEvent(db)

	event := entity.Event{ID: 1, Name: "Teste", StartTime: time.Now(), EndTime: time.Now().AddDate(0, 0, 1), Description: "Test"}
	err := eventDB.Update(&event)
	assert.Nil(t, err)
}

func TestDeleteById(t *testing.T) {
	configs.ConnectGorm()
	db := configs.GetDB()
	migrations.MigrateTable()
	eventDB := NewEvent(db)

	err := eventDB.Delete("1")
	assert.Nil(t, err)
	err = eventDB.Delete("0")
	assert.Error(t, err)

}
