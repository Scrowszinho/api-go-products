package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Scrowszinho/api-go-products/internal/dto"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/Scrowszinho/api-go-products/internal/infra/database"
)

type BetsHandler struct {
	BetsDB database.BetsInterface
}

func NewBetsHandler(db database.BetsInterface) *BetsHandler {
	return &BetsHandler{
		BetsDB: db,
	}
}

func (e *BetsHandler) CreateBets(w http.ResponseWriter, r *http.Request) {
	var bets dto.CreateBetsInput
	err := json.NewDecoder(r.Body).Decode(&bets)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	b, err := entity.NewBet(&bets.User, bets.OutcomeID, bets.Amount, bets.Bonus, bets.Active)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = e.BetsDB.Create(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
