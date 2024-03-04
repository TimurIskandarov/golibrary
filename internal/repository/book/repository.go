package repository

import (
	"context"
	"errors"

	"golibrary/internal/model"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

const (
	books = "books"
)

type BookerRepository interface {
	BookTake(ctx context.Context, userId, bookId int) error
	BookReturn(ctx context.Context, userId, bookId int) error
	BooksList(ctx context.Context) ([]*model.Book, error)
	BookAdd(ctx context.Context, book model.Book) error
}

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) BookTake(ctx context.Context, userId, bookId int) error {
	var isAvailable bool
	query, args, _ := sq.Select("available").
		From(books).
		Where("id=?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&isAvailable); err != nil {
		return err
	}

	if !isAvailable {
		return errors.New("книга недоступна")
	}

	query, args, _ = sq.Update(books).
		Set("available", false).
		Set("user_id", userId).
		Where("id = ?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	// TODO: добавить запись о взятии книги в историю
	// ...

	return nil
}

func (r *BookRepository) BookReturn(ctx context.Context, userId, bookId int) error {
	var isAvailable bool
	query, args, _ := sq.Select("available").
		From(books).
		Where("id=?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&isAvailable); err != nil {
		return err
	}

	if isAvailable {
		return errors.New("книга уже сдана")
	}

	query, args, _ = sq.Update(books).
		Set("available", true).
		Set("user_id", nil).
		Where("id = ?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	// TODO: удалить запись о взятии книги из истории
	// ...

	return nil
}

func (r *BookRepository) BooksList(ctx context.Context) ([]*model.Book, error) {
	query, args, _ := sq.Select(
		"id",
		"title",
		"available",
		"user_id",
		"author_id",
	).From(
		books,
	).PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := make([]*model.Book, 0)

	for rows.Next() {
		book := new(model.Book)

		if err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Available,
			&book.UserID,
			&book.AuthorID,
		); err != nil {
			return nil, err
		}

		// уникальные книги
		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, errors.New("в наличии книг нет")
	}

	// убеждаемся, что прошлись по всему набору строк без ошибок
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) BookAdd(ctx context.Context, book model.Book) error {
	var bookId int
	query, args, _ := sq.Select("id").
		From(books).
		Where("title=?", book.Title).
		Where("author_id=?", book.AuthorID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&bookId); err != nil {
		return err
	}

	if bookId == 0 {
		query, args, _ := sq.Update(books).
			Set("title", book.Title).
			Set("available", book.Available).
			Set("author_id", book.AuthorID).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
			return err
		}
	}

	return nil
}
