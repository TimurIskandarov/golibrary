package db

import (
	"database/sql"
	"fmt"
	"time"

	"golibrary/config"
	"golibrary/internal/db/migrator"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

const (
	numUsers   = 57
	numAuthors = 10
	numBooks   = 10
	users      = "users"
	authors    = "authors"
	books      = "books"
)

func checkLibrary(db *sqlx.DB, table string) bool {
	var rows int
	query, _, _ := sq.Select("*").From(table).ToSql()
	db.QueryRow(query).Scan(&rows)
	return rows == 0
}

func addRandomUsers(db *sqlx.DB) error {
	for i := 0; i < numUsers; i++ {
		name := gofakeit.Name()
		email := gofakeit.Email()

		query, args, _ := sq.Insert(users).
			Columns(
				"name",
				"email",
			).Values(
			name,
			email,
		).PlaceholderFormat(sq.Dollar).ToSql()

		_, err := db.Exec(query, args...)

		if err != nil {
			return err
		}
	}
	return nil
}

func addAuthors(db *sqlx.DB) error {
	for i := 0; i < numAuthors; i++ {
		var authorId int
		authorName := gofakeit.Name()
		birthDate := fmt.Sprint(gofakeit.Date().Format("2006-01-02"))

		query, args, _ := sq.Insert(authors).
			Columns(
				"name",
				"birth_date",
			).Values(
			authorName,
			birthDate,
		).Suffix(
			"RETURNING id",
		).PlaceholderFormat(sq.Dollar).ToSql()

		row := db.QueryRow(query, args...)
		err := row.Scan(&authorId)

		if err != nil {
			return err
		}

		err = addBooks(db, authorId)
		if err != nil {
			return err
		}
	}

	return nil
}

func addBooks(db *sqlx.DB, authorId int) error {
	var bookId int

	for i := 0; i < numBooks; i++ {
		bookTitle := gofakeit.BookTitle()

		query, args, _ := sq.Insert(books).
			Columns(
				"title",
				"available",
				"author_id",
			).Values(
				bookTitle,
				true,
				authorId,
			).Suffix(
				"RETURNING id",
			).PlaceholderFormat(sq.Dollar).ToSql()

		row := db.QueryRow(query, args...)
		err := row.Scan(&bookId)
		if err != nil {
			return err
		}

		// TODO: добавить данные в связь автор_книга
	}

	return nil
}

func NewSqlDB(dbRaw *sql.DB, driver string, logger *zap.Logger) (*sqlx.DB, error) {
	db := sqlx.NewDb(dbRaw, driver)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)

	migrator.CreateTables(db)

	usersNotFound := checkLibrary(db, users)
	if usersNotFound {
		err := addRandomUsers(db)
		if err != nil {
			logger.Error("error add user", zap.Error(err))
		}
	}

	authorsNotFound, booksNotFound := checkLibrary(db, authors), checkLibrary(db, books)
	if authorsNotFound && booksNotFound {
		err := addAuthors(db)
		if err != nil {
			logger.Error("error authors and books", zap.Error(err))
		}
	}

	return db, nil

}

func Init(dbConf config.DB, logger *zap.Logger) (*sqlx.DB, error) {
	var dsn string
	var err error
	var dbRaw *sql.DB

	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * time.Duration(dbConf.Timeout))

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %d timeout %s", dbConf.Timeout, err)
		case <-ticker.C:
			dbRaw, err = sql.Open(dbConf.Driver, dsn)
			if err != nil {
				return nil, err
			}
			err = dbRaw.Ping()
			if err == nil {
				return NewSqlDB(dbRaw, dbConf.Driver, logger)
			}
			logger.Error("failed to connect to the database", zap.String("dsn", dsn), zap.Error(err))
		}
	}
}
