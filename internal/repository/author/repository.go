package repo

import (
	"context"

	"golibrary/internal/model"
	repoBook "golibrary/internal/repository/book"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

const (
	authors = "authors"
)

type AuthorerRepository interface {
	AuthorsTop(ctx context.Context) ([]*model.Author, error)
	AuthorsList(ctx context.Context) ([]*model.Author, error)
	AddAuthor(ctx context.Context, name, birthDate string) error
}

type AuthorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}
// +
func (r *AuthorRepository) AuthorsList(ctx context.Context) ([]*model.Author, error) {
	query, _, _ := sq.Select(
		"id",
		"name",
		"birth_date",
	).From(
		authors,
	).PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authors := make([]*model.Author, 0)

	for rows.Next() {
		author := new(model.Author)

		if err = rows.Scan(
			&author.ID,
			&author.Name,
			&author.BirthDate,
		); err != nil {
			return nil, err
		}

		repoBook := repoBook.NewBookRepository(r.db)
		author.Books, err = repoBook.ListByAuthor(ctx, author.ID)
		if err != nil {
			return nil, err // authors, err
		}

		authors = append(authors, author)
	}

	// убеждаемся, что прошлись по всему набору строк без ошибок
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *AuthorRepository) AuthorsTop(ctx context.Context) ([]*model.Author, error) {
	query, _, _ := sq.Select(
		"id",
		"name",
	).From(
		authors,
	).PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authors := make([]*model.Author, 0)

	for rows.Next() {
		author := new(model.Author)

		if err = rows.Scan(
			&author.ID,
			&author.Name,
			&author.BirthDate,
		); err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	// убеждаемся, что прошлись по всему набору строк без ошибок
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *AuthorRepository) AddAuthor(ctx context.Context, name, birthDate string) error {
	var authorId int
	query, args, _ := sq.Insert(authors).
		Columns(
			"name",
			"birth_date",
		).Values(
		name,
		birthDate,
	).Suffix(
		"RETURNING id",
	).PlaceholderFormat(sq.Dollar).ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&authorId); err != nil {
		return err
	}

	return nil
}
