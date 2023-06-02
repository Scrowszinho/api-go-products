package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Scrowszinho/api-go-products/internal/dto"
	"github.com/Scrowszinho/api-go-products/internal/entity"
	"github.com/Scrowszinho/api-go-products/internal/infra/database"
	"github.com/go-chi/chi"
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
	id, ok := r.Context().Value("user").(int)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	b, err := entity.NewBet(id, bets.OutcomeID, bets.Amount, bets.Bonus, bets.Active)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = e.BetsDB.Create(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (e *BetsHandler) GetBet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bets, err := e.BetsDB.FindById(idt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bets)
}

func (e *BetsHandler) DeleteBets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = e.BetsDB.FindById(idt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = e.BetsDB.Delete(idt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (e *BetsHandler) UpdateBets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var bets entity.Bets
	err := json.NewDecoder(r.Body).Decode(&bets)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = e.BetsDB.FindById(bets.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = e.BetsDB.Update(&bets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
