package repo

import (
	"context"

	"golibrary/internal/model"
	repoBook "golibrary/internal/repository/book"

	"github.com/jmoiron/sqlx"
	
	sq "github.com/Masterminds/squirrel"
)

const (
	users = "users"
)

type Userer interface {
	List(ctx context.Context) ([]*model.User, error)
	Add(ctx context.Context, book model.User) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) List(ctx context.Context) ([]*model.User, error) {
	query, args, _ := sq.Select(
			"id",
			"name",
			"email",
		).
		From(users).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*model.User, 0)

	for rows.Next() {
		user := new(model.User)
		
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		); err != nil {
			return nil, err
		}
		
		repoBook := repoBook.NewBookRepository(r.db)
		user.Books, err = repoBook.ListByUser(ctx, user.ID)
		if err != nil {
			return nil, err // users, err
		}

		users = append(users, user)
	}

	// убеждаемся, что прошлись по всему набору строк без ошибок
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Add(ctx context.Context, book model.User) error {
	return nil
}
