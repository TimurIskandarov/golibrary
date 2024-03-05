package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Books []Book `json:"books,omitempty"`
}

type Author struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	BirthDate string  `json:"birth_date"`
	Books     []*Book `json:"books"`
}

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Available bool   `json:"available"`
	UserID    int    `json:"user_id,omitempty"`
	AuthorID  int    `json:"author_id"`
}
