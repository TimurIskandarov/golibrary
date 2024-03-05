package controller

import (
	"golibrary/internal/service"

	acontroller "golibrary/internal/controller/author"
	bcontroller "golibrary/internal/controller/book"
	ucontroller "golibrary/internal/controller/user"
)

type Controllers struct {
	Author acontroller.Authorer
	Book   bcontroller.Booker
	User   ucontroller.Userer
}

func NewControllers(services *service.Services) *Controllers{
	authorController := acontroller.NewAuthor(services.Author)
	bookController := bcontroller.NewBook(services.Book)
	userController := ucontroller.NewUser(services.User)

	return &Controllers{
		Author: authorController,
		Book:   bookController,
		User:   userController,
	}
}
