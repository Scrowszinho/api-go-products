package main

import (
	"net/http"

	"github.com/Scrowszinho/api-go-products/configs"
	"github.com/Scrowszinho/api-go-products/internal/infra/database"
	"github.com/Scrowszinho/api-go-products/internal/webserver/handlers"
	"github.com/Scrowszinho/api-go-products/migrations"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
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
	eventDB := database.NewEvent(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)
	eventHanler := handlers.NewEventHandler(eventDB)
	userRoutes(r, *userHandler)
	eventsRoutes(r, *eventHanler, config)
	http.ListenAndServe(":8000", r)

}

func userRoutes(r *chi.Mux, userHandler handlers.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/login", userHandler.GetJWT)
	})
}

func eventsRoutes(r *chi.Mux, eventsHandler handlers.EventHandler, config *configs.Configs) {
	r.Route("/events", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", eventsHandler.CreateEvent)
		r.Get("/{id}", eventsHandler.GetEvent)
		r.Get("/", eventsHandler.GetEvents)
		r.Put("/{id}", eventsHandler.UpdateEvent)
		r.Delete("/{id}", eventsHandler.DeleteEvent)

	})
}
