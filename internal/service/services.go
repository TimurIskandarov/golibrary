package service

import (
	"golibrary/internal/repository"
	aservice "golibrary/internal/service/author"
	bservice "golibrary/internal/service/book"
	uservice "golibrary/internal/service/user"

	"go.uber.org/zap"
)

type Services struct {
	Author aservice.Authorer
	Book   bservice.Booker
	User   uservice.Userer
}

func NewServices(repos *repo.Repositories, logger *zap.Logger) *Services {
	return &Services{
		Author: aservice.NewAuthorService(repos.Author, logger),
		Book:   bservice.NewBookService(repos.Book, logger),
		User:   uservice.NewUserService(repos.User, logger),
	}
}
