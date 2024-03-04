package controller

import (
	"encoding/json"
	"net/http"

	"golibrary/internal/model"
	"golibrary/internal/service/author"
)

type Authorer interface {
	Add(w http.ResponseWriter, r *http.Request)
	Top(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type Author struct {
	service service.Authorer
}

func NewAuthor(service service.Authorer) Authorer {
	return &Author{
		service: service,
	}
}

func (a *Author) Top(w http.ResponseWriter, r *http.Request) {
	authors, err := a.service.AuthorsTop(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authors)
}

func (a *Author) List(w http.ResponseWriter, r *http.Request) {
	authors, err := a.service.AuthorsList(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authors)
}

func (a *Author) Add(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.service.AddAuthor(r.Context(), author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(author)
}
