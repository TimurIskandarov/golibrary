package repo

import (
	"context"
	"database/sql"
	"errors"

	"golibrary/internal/model"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

const (
	books = "books"
)

type Booker interface {
	Take(ctx context.Context, userId, bookId int) (*model.Book, error)
	Return(ctx context.Context, userId, bookId int) (*model.Book, error)
	List(ctx context.Context) ([]*model.Book, error)
	Add(ctx context.Context, book model.Book) (int, error)
	ListByAuthor(ctx context.Context, authorId int) ([]*model.Book, error)
	ListByUser(ctx context.Context, authorId int) ([]*model.Book, error)
	ReadableAuthors(ctx context.Context) ([]int, error)
}

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Take(ctx context.Context, userId, bookId int) (*model.Book, error) {
	user := new(model.User)
	query, args, _ := sq.Select("id").
		From("users").
		Where("id=?", userId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&user.ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("такого читателя нет")
		}
		return nil, err
	}

	var userID *int
	book := new(model.Book)
	query, args, _ = sq.Select(
			"id",
			"title",
			"user_id",
			"author_id",
		).
		From(books).
		Where("id=?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row = r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(
		&book.ID,
		&book.Title,
		&userID,
		&book.AuthorID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("такой книги нет")
		}
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

func (r *BookRepository) Return(ctx context.Context, userId, bookId int) (*model.Book, error) {
	user := new(model.User)
	query, args, _ := sq.Select("id").
		From("users").
		Where("id=?", userId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&user.ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("такого читателя нет")
		}
		return nil, err
	}
	
	var userID *int
	book := new(model.Book)
	query, args, _ = sq.Select(
			"id",
			"title",
			"user_id",
			"author_id",
		).
		From(books).
		Where("id=?", bookId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row = r.db.QueryRowContext(ctx, query, args...)
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

func (r *BookRepository) List(ctx context.Context) ([]*model.Book, error) {
	query, args, _ := sq.Select(
			"id",
			"title",
			"user_id",
			"author_id",
		).
		From(books).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := make([]*model.Book, 0)

	for rows.Next() {
		var userId *int
		book := new(model.Book)

		if err = rows.Scan(
			&book.ID,
			&book.Title,
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

	if len(books) == 0 {
		return nil, errors.New("в наличии книг нет")
	}

	// убеждаемся, что прошлись по всему набору строк без ошибок
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Add(ctx context.Context, book model.Book) (int, error) {
	var authorId int
	query, args, _ := sq.Select("id").
		From("authors").
		Where("id=?", book.AuthorID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&authorId); err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("такого автора не существует")
		}
	}

	// поиск книги по автору и названию
	var bookId int
	query, args, _ = sq.Select("id").
		From(books).
		Where("title=?", book.Title).
		Where("author_id=?", book.AuthorID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row = r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&bookId); err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	}

	// если книги данного автора нет, то добавляем её
	if bookId == 0 {
		query, args, _ = sq.Insert(books).
			Columns(
				"title",
				"author_id",
			).
			Values(
				book.Title,
				book.AuthorID,
			).
			Suffix("RETURNING id").
			PlaceholderFormat(sq.Dollar).
			ToSql()

		row := r.db.QueryRowContext(ctx, query, args...)
		if err := row.Scan(&bookId); err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("такая книга уже есть")
	}

	return bookId, nil
}

func (r *BookRepository) ListByAuthor(ctx context.Context, authorId int) ([]*model.Book, error) {
	query, args, _ := sq.Select(
			"id",
			"title",
			"user_id",
			"author_id",
		).
		From(books).
		Where("author_id=?", authorId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

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

func (r *BookRepository) ListByUser(ctx context.Context, userId int) ([]*model.Book, error) {
	query, args, _ := sq.Select(
			"id",
			"title",
			"user_id",
			"author_id",
		).
		From(books).
		Where("user_id=?", userId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := make([]*model.Book, 0)

	for rows.Next() {
		book := new(model.Book)

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.UserID,
			&book.AuthorID,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) ReadableAuthors(ctx context.Context) ([]int, error) {
	query, _, _ := sq.Select(
			"author_id",
			"count(author_id) as rating",
		).
		From(books).
		Where(sq.NotEq{"user_id": nil}).
		GroupBy("author_id").
		OrderBy("rating DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authorsIds := make([]int, 0)

	for rows.Next() {
		var authorId, rating int
		if err = rows.Scan(
			&authorId,
			&rating,
		); err != nil {
			return nil, err
		}

		authorsIds = append(authorsIds, authorId)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authorsIds, nil
}
