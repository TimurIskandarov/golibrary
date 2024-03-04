package repo

import (
	arepo "golibrary/internal/repository/author"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Author arepo.AuthorerRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Author: arepo.NewAuthorRepository(db),
	}
}
