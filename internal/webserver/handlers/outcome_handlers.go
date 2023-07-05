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

type OutcomeHandler struct {
	OutcomeDB database.OutcomeInterface
}

func NewOutcomeHandler(db database.OutcomeInterface) *OutcomeHandler {
	return &OutcomeHandler{
		OutcomeDB: db,
	}
}

func (o *OutcomeHandler) CreateOutcome(w http.ResponseWriter, r *http.Request) {
	var outcome dto.CreateOutcomesInput
	err := json.NewDecoder(r.Body).Decode(&outcome)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.CreateOutcome(outcome.Name, outcome.Odds, outcome.Status, outcome.MarketID, outcome.BetID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = o.OutcomeDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(outcome)
	w.WriteHeader(http.StatusCreated)

}

func (o *OutcomeHandler) GetOutcome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idR, _ := strconv.Atoi(id)
	outcome, err := o.OutcomeDB.FindById(idR)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outcome)
}

func (o *OutcomeHandler) UpdateOutcome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var outcome entity.Outcome
	err := json.NewDecoder(r.Body).Decode(&outcome)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idR, _ := strconv.Atoi(id)
	outcome.ID = idR
	_, err = o.OutcomeDB.FindById(idR)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = o.OutcomeDB.Update(&outcome)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (o *OutcomeHandler) DeleteOutcome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idR, _ := strconv.Atoi(id)
	_, err := o.OutcomeDB.FindById(idR)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = o.OutcomeDB.Delete(idR)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (o *OutcomeHandler) GetOutcomesByEventId(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sort := r.URL.Query().Get("sort")
	products, err := o.OutcomeDB.FindOutcomesById(id, page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
