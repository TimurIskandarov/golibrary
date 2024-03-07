package docs

import "golibrary/internal/model"

// swagger:route PUT /library/take/{bookId}/{userId} book BookTake
//		Взять книгу
// security:
// - basic
// responses:
// 	 200: body:Book Книга успешно получена

// swagger:model Book
type BookTake struct {
	// in:body
	Book Book
}

// swagger:route PUT /library/return/{bookId}/{userId} book BookReturn
//		Сдать книгу
// security:
// - basic
// responses:
// 	 200: body:Book Книга успешно сдана

// swagger:model Book
type BookReturn struct {
	// in:body
	Book Book
}

// swagger:route GET /library/books book BooksList
//		Список книг
// security:
// - basic
// responses:
// 	 200: body:Book Список книг успешно получен

// swagger:model Books
type Books []model.Book

// swagger:route POST /library/book book BookAdd
//		Добавить книгу
// security:
// - basic
// responses:
// 	 200: body:Book Книга успешно добавлена

// swagger:model Book
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	UserID   int    `json:"userId"`
	AuthorID int    `json:"authorId"`
}
