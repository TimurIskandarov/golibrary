package repo

import (
	arepo "golibrary/internal/repository/author"
	brepo "golibrary/internal/repository/book"
	urepo "golibrary/internal/repository/user"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Author arepo.AuthorerRepository
	Book   brepo.BookerRepository
	User   urepo.Userer
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Author: arepo.NewAuthorRepository(db),
		Book:   brepo.NewBookRepository(db),
		User:   urepo.NewUserRepository(db),
	}
}
