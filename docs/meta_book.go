package docs

import "golibrary/internal/model"

// swagger:route PUT /library/take/{bookId}/{userId} book BookTake
//		Взять книгу
// Security:
// - basic
// Responses:
// 	 200: body:Book Книга успешно получена

// swagger:parameters BookTake
type _ struct {
	//in:path
	BookId int `json:"bookId"`
	//in:path
	UserId int `json:"userId"`
}

// swagger:model Book
type BookTake struct {
	// in:body
	ID       int    `json:"id"`
	Title    string `json:"title"`
	UserID   int    `json:"userId"`
	AuthorID int    `json:"authorId"`
}

// swagger:route PUT /library/return/{bookId}/{userId} book BookReturn
//		Сдать книгу
// Security:
// - basic
// Responses:
// 	 200: body:Book Книга успешно сдана

// swagger:parameters BookReturn
type _ struct {
	//in:path
	BookId int `json:"bookId"`
	//in:path
	UserId int `json:"userId"`
}

// swagger:model Book
type BookReturn struct {
	// in:body
	ID       int    `json:"id"`
	Title    string `json:"title"`
	UserID   int    `json:"userId"`
	AuthorID int    `json:"authorId"`
}

// swagger:route GET /library/books book BooksList
//		Список книг
// Security:
// - basic
// Responses:
// 	 200: body:Book Список книг успешно получен

// swagger:parameters BooksList

// swagger:model Books
type Books []model.Book
