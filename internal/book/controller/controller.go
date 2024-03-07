package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golibrary/internal/model"
	"golibrary/internal/book/service"

	"github.com/go-chi/chi"
)

type Booker interface {
	Add(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Take(w http.ResponseWriter, r *http.Request)
	Return(w http.ResponseWriter, r *http.Request)
}

type Book struct {
	service service.Booker
}

func NewBook(service service.Booker) Booker {
	return &Book{
		service: service,
	}
}

// Take implements Booker.
func (b *Book) Take(w http.ResponseWriter, r *http.Request) {
	rawUserId := chi.URLParam(r, "userId")
	rawBookId := chi.URLParam(r, "bookId")

	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookId, err := strconv.Atoi(rawBookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := b.service.Take(r.Context(), userId, bookId)
	if err != nil {
		if err.Error() == "книга недоступна" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Return implements Booker.
func (b *Book) Return(w http.ResponseWriter, r *http.Request) {
	rawUserId := chi.URLParam(r, "userId")
	rawBookId := chi.URLParam(r, "bookId")

	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookId, err := strconv.Atoi(rawBookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := b.service.Return(r.Context(), userId, bookId)
	if err != nil {
		if err.Error() == "книга уже сдана" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// List implements Booker.
func (b *Book) List(w http.ResponseWriter, r *http.Request) {
	books, err := b.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Add implements Booker.
func (b *Book) Add(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID, err = b.service.Add(r.Context(), book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
