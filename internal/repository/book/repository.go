package repo

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
	BookTake(ctx context.Context, userId, bookId int) (*model.Book, error)
	BookReturn(ctx context.Context, userId, bookId int) (*model.Book, error)
	BooksList(ctx context.Context) ([]*model.Book, error)
	BookAdd(ctx context.Context, book model.Book) error
	ListByAuthor(ctx context.Context, authorId int) ([]*model.Book, error)
}

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) BookTake(ctx context.Context, userId, bookId int) (*model.Book, error) {
	var userID *int
	book := new(model.Book)
	query, args, _ := sq.Select(
			"id",
			"title",
			"user_id",
			"author_id",
		).
		From(books).
		Where("id=?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(
		&book.ID,
		&book.Title,
		&userID,
		&book.AuthorID,
	); err != nil {
		return nil, err
	}

	if userID != nil {
		return nil, errors.New("книга недоступна")
	} else {
		book.UserID = userId
	}

	query, args, _ = sq.Update(books).
		Set("user_id", book.UserID).
		Where("id = ?", book.ID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return nil, err // book, err
	}

	return book, nil
}

func (r *BookRepository) BookReturn(ctx context.Context, userId, bookId int) (*model.Book, error) {
	var userID *int
	book := new(model.Book)
	query, args, _ := sq.Select(
			"id",
			"title",
			"user_id",
			"author_id",
		).
		From(books).
		Where("id=?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(
		&book.ID,
		&book.Title,
		&userID,
		&book.AuthorID,
	); err != nil {
		return nil, err
	}

	if userID == nil {
		return nil, errors.New("книга уже сдана")
	} else if *userID != userId {
		return nil, errors.New("вы книгу не брали")
	} else {
		userID = nil
	}

	query, args, _ = sq.Update(books).
		Set("user_id", userID).
		Where("id = ?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return nil, err // book, err
	}

	return book, nil
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

// +
func (r *BookRepository) ListByAuthor(ctx context.Context, authorId int) ([]*model.Book, error) {
	query, args, _ := sq.Select(
		"id",
		"title",
		"available",
		"user_id",
		"author_id",
	).From(
		books,
	).Where(
		"author_id = ?", authorId,
	).PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := make([]*model.Book, 0)

	for rows.Next() {
		book := new(model.Book)
		var userId *int

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Available,
			&userId,
			&book.AuthorID,
		); err != nil {
			return nil, err
		}

		if userId == nil {
			book.UserID = 0
		} else {
			book.UserID = *userId
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
