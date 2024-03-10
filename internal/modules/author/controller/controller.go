package controller

import (
	"encoding/json"
	"net/http"

	"golibrary/internal/modules/author/repository"
	"golibrary/internal/modules/author/service"
	"golibrary/internal/model"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Authorer interface {
	Top(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}

type Author struct {
	service service.Authorer
}

func NewAuthor(db *sqlx.DB, logger *zap.Logger) Authorer {
	return &Author{
		service: service.NewAuthorService(
			repo.NewAuthorRepository(db), 
			logger,
		),
	}
}

func (a *Author) Top(w http.ResponseWriter, r *http.Request) {
	authors, err := a.service.Top(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (a *Author) List(w http.ResponseWriter, r *http.Request) {
	authors, err := a.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (a *Author) Add(w http.ResponseWriter, r *http.Request) {
	var author model.Author

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.service.Add(r.Context(), author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}
