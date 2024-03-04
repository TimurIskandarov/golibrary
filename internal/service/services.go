package service

import (
	aservice "golibrary/internal/service/author"
	bservice "golibrary/internal/service/book"

	"go.uber.org/zap"
)

type Services struct {
	Author aservice.Authorer
	Book   bservice.Booker
}

func NewServices(logger *zap.Logger) *Services {
	return &Services{
		Author: aservice.NewAuthorService(logger),
		Book: bservice.NewBookService(logger),
	}
}
