package migrations

import (
	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}, &entity.Events{}, &entity.Outcome{}, &entity.Bets{}}
}

func MigrateTable() {
	db := configs.GetDB()
	db.AutoMigrate(getModels()...)
}
