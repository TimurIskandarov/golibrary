package facade

import (
	"golibrary/internal/superservice"
	"net/http"
)

type Facade struct {
	Service *superservice.Service
}

func (f *Facade) UsersList(w http.ResponseWriter, r *http.Request) {
	f.Service.User.List(w, r)
}

func (f *Facade) UserAdd(w http.ResponseWriter, r *http.Request) {
	f.Service.User.Add(w, r)
}

func (f *Facade) AuthorsTop(w http.ResponseWriter, r *http.Request) {
	f.Service.Author.Top(w, r)
}

func (f *Facade) AuthorsList(w http.ResponseWriter, r *http.Request) {
	f.Service.Author.List(w, r)
}

func (f *Facade) AuthorAdd(w http.ResponseWriter, r *http.Request) {
	f.Service.Author.Add(w, r)
}

func (f *Facade) BookTake(w http.ResponseWriter, r *http.Request) {
	f.Service.Book.Take(w, r)
}

func (f *Facade) BookReturn(w http.ResponseWriter, r *http.Request) {
	f.Service.Book.Return(w, r)
}

func (f *Facade) BookList(w http.ResponseWriter, r *http.Request) {
	f.Service.Book.List(w, r)
}

func (f *Facade) BookAdd(w http.ResponseWriter, r *http.Request) {
	f.Service.Book.Add(w, r)
}
