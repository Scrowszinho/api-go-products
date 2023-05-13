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
