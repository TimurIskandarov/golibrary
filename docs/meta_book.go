package docs

// swagger:route PUT /library/take/{bookId}/{userId} book BookTake
//		Взять книгу
// responses:
// 	 200: body:Book Книга успешно получена

// swagger:route PUT /library/return/{bookId}/{userId} book BookReturn
//		Сдать книгу
// responses:
// 	 200: body:Book Книга успешно сдана

// swagger:route GET /library/books book BooksList
//		Список книг
// responses:
// 	 200: body:[]Book Список книг успешно получен

// swagger:route POST /library/book book BookAdd
//		Добавить книгу
// responses:
// 	 200: body:Book Книга успешно добавлена

// swagger:model Book
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	UserID   int    `json:"userId"`
	AuthorID int    `json:"authorId"`
}
