package superservice

import (
	"net/http"

	author "golibrary/internal/modules/author/controller"
	book "golibrary/internal/modules/book/controller"
	user "golibrary/internal/modules/user/controller"
)

type Service struct {
	Author author.Authorer
	Book   book.Booker
	User   user.Userer
}

func (s *Service) UsersList(w http.ResponseWriter, r *http.Request) {
	s.User.List(w, r)
}

func (s *Service) UserAdd(w http.ResponseWriter, r *http.Request) {
	s.User.Add(w, r)
}

func (s *Service) AuthorsTop(w http.ResponseWriter, r *http.Request) {
	s.Author.Top(w, r)
}

func (s *Service) AuthorsList(w http.ResponseWriter, r *http.Request) {
	s.Author.List(w, r)
}

func (s *Service) AuthorAdd(w http.ResponseWriter, r *http.Request) {
	s.Author.Add(w, r)
}

func (s *Service) BookTake(w http.ResponseWriter, r *http.Request) {
	s.Book.Take(w, r)
}

func (s *Service) BookReturn(w http.ResponseWriter, r *http.Request) {
	s.Book.Return(w, r)
}

func (s *Service) List(w http.ResponseWriter, r *http.Request) {
	s.Book.List(w, r)
}

func (s *Service) BookAdd(w http.ResponseWriter, r *http.Request) {
	s.Book.Add(w, r)
}
