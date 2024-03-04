package migrator

import "github.com/jmoiron/sqlx"

// временное решение
func CreateTables(db *sqlx.DB) {
	CreateAuthors(db)
	CreateBooks(db)
	CreateUsers(db)
	CreateHistory(db)
}

func CreateAuthors(db *sqlx.DB) error {
	query := `CREATE TABLE public.authors (
		id serial PRIMARY KEY,
		name VARCHAR NOT NULL,
		birth_date VARCHAR(10) NOT NULL
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
		available BOOLEAN DEFAULT true,
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
		email VARCHAR
	);`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}

func CreateHistory(db *sqlx.DB) error {
	query := `CREATE TABLE public.history (
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		book_id INT REFERENCES books(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, book_id)
	);`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
