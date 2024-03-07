package repo

import (
	"context"

	"golibrary/internal/model"
	repoBook "golibrary/internal/book/repository"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

const (
	authors = "authors"
)

type Authorer interface {
	Top(ctx context.Context) ([]*model.Author, error)
	List(ctx context.Context) ([]*model.Author, error)
	Add(ctx context.Context, author model.Author) error
}

type AuthorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) List(ctx context.Context) ([]*model.Author, error) {
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

func (r *AuthorRepository) Top(ctx context.Context) ([]*model.Author, error) {
	repoBook := repoBook.NewBookRepository(r.db)
	authorsIds, err := repoBook.ReadableAuthors(ctx)
	if err != nil {
		return nil, nil
	}

	query, args, _ := sq.Select(
			"id",
			"name",
			"birth_date",
		).
		From(authors).
		Where(sq.Eq{"id": authorsIds}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authors := make(map[int]*model.Author, len(authorsIds))

	for rows.Next() {
		author := new(model.Author)

		if err = rows.Scan(
			&author.ID,
			&author.Name,
			&author.BirthDate,
		); err != nil {
			return nil, err
		}

		author.Books, err = repoBook.ListByAuthor(ctx, author.ID)
		if err != nil {
			return nil, err // authors, err
		}

		authors[author.ID] = author
	}

	// убеждаемся, что прошлись по всему набору строк без ошибок
	if err := rows.Err(); err != nil {
		return nil, err
	}

	authorsTop := make([]*model.Author, len(authorsIds))
	for i, authorId := range authorsIds {
		authorsTop[i] = authors[authorId]
	}

	return authorsTop, nil
}

func (r *AuthorRepository) Add(ctx context.Context, author model.Author) error {
	var authorId int
	query, args, _ := sq.Insert(authors).
		Columns(
			"name",
			"birth_date",
		).
		Values(
			author.Name,
			author.BirthDate,
		).
		Suffix(
			"ON CONFLICT (name) DO UPDATE SET name = ? RETURNING id", 
			author.Name,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&authorId); err != nil {
		return err
	}

	repoBook := repoBook.NewBookRepository(r.db)
	for _, book := range author.Books {
		book.AuthorID = authorId
		repoBook.Add(ctx, *book)
	}

	return nil
}
