package controller

import (
	"golibrary/internal/service"

	acontroller "golibrary/internal/controller/author"
	bcontroller "golibrary/internal/controller/book"
)

type Controllers struct {
	Author acontroller.Authorer
	Book   bcontroller.Booker
}

func NewControllers(services *service.Services) *Controllers{
	authorController := acontroller.NewAuthor(services.Author)
	bookController := bcontroller.NewBook(services.Book)

	return &Controllers{
		Author: authorController,
		Book:   bookController,
	}
}
