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
	Add(ctx context.Context, book model.User) (int, error)
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

func (r *UserRepository) Add(ctx context.Context, user model.User) (int, error) {
	var userId int
	var err error
	query, args, _ := sq.Insert(users).
		Columns(
			"name",
			"email",
		).
		Values(
			user.Name,
			user.Email,
		).
		Suffix(
			"ON CONFLICT (email) DO UPDATE SET email = ? RETURNING id", 
			user.Email,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&userId); err != nil {
		return 0, err
	}

	repoBook := repoBook.NewBookRepository(r.db)
	for i, book := range user.Books {
		book.UserID = userId
		user.Books[i], err = repoBook.Take(ctx, userId, book.ID)
		if err != nil {
			return userId, err
		}
	}

	return userId, nil
}
