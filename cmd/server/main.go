package main

import (
	"fmt"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.GetConfigs(".")
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{})
}
