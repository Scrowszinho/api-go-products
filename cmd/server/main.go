package main

import (
	"net/http"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/infra/database"
	"github.com/Scrowszinho/api-go-products/internal/webserver/handlers"
	"github.com/Scrowszinho/api-go-products/migrations"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	configs.ConnectGorm()
	db := configs.GetDB()
	db.AutoMigrate()
	migrations.MigrateTable()

	config, err := configs.GetConfigs(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/login", userHandler.GetJWT)
	http.ListenAndServe(":8000", r)

}
