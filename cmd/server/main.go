package main

import (
	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/migrations"
)

func main() {
	configs.ConnectGorm()
	db := configs.GetDB()
	db.AutoMigrate()
	migrations.MigrateTable()
}
