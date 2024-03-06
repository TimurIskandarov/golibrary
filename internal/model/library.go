package model

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Books []*Book `json:"books,omitempty"`
}

type Author struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	BirthDate string  `json:"birthDate"`
	Books     []*Book `json:"books"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	UserID   int    `json:"userId,omitempty"`
	AuthorID int    `json:"authorId"`
}
