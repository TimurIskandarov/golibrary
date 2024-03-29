package migrator

import "github.com/jmoiron/sqlx"

// временное решение
func CreateTables(db *sqlx.DB) {
	CreateAuthors(db)
	CreateBooks(db)
	CreateUsers(db)
}

func CreateAuthors(db *sqlx.DB) error {
	query := `CREATE TABLE public.authors (
		id serial PRIMARY KEY,
		name VARCHAR NOT NULL,
		birth_date VARCHAR(10) NOT NULL,
		CONSTRAINT name_key UNIQUE (name)
	);`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}

func CreateBooks(db *sqlx.DB) error {
	query := `CREATE TABLE public.books (
		id serial PRIMARY KEY,
		title VARCHAR NOT NULL,
		user_id INT,
		author_id INT NOT NULL,
		CONSTRAINT fk_books_1 FOREIGN KEY (author_id) REFERENCES public.authors(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}

func CreateUsers(db *sqlx.DB) error {
	query := `CREATE TABLE public.users (
		id serial PRIMARY KEY,
		name VARCHAR,
		email VARCHAR,
		CONSTRAINT email_key UNIQUE (email)
	);`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
