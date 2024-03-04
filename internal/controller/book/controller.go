package controller

import (
	"net/http"
	"strconv"

	service "golibrary/internal/service/book"

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

// AddBook implements Booker.
func (b *Book) Add(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// BookReturn implements Booker.
func (b *Book) Return(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// BookTake implements Booker.
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

	b.service.BookTake(r.Context(), userId, bookId)
}

// ListBooks implements Booker.
func (b *Book) List(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
