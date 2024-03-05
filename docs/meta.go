// Package classification Golibrary API.
//
// Документация Golibrary API.
//
//		Schemes: http, https
//		Host: localhost:8080
//		BasePath: /
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//	    - application/x-www-form-urlencoded
//		- multipart/form-data
//
//		Produces:
//		- application/json
//
//		Security:
//		- api_key:
//
//
//		SecurityDefinitions:
//		  api_key:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//
// swagger:meta
package docs

import "golibrary/internal/model"

// swagger:route GET /library/authors author Author
//		Получение списка авторов
// security:
// - basic
// responses:
// 	 200: body:Author Cписок авторов успешно получен

// swagger:parameters Author

// swagger:model Author
type Author struct {
	// in:body
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	BirthDate string        `json:"birth_date"`
	Books     []*model.Book `json:"books"`
}

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
	UserID   int    `json:"user_id"`
	AuthorID int    `json:"author_id"`
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
	UserID   int    `json:"user_id"`
	AuthorID int    `json:"author_id"`
}

// swagger:route GET /library/users user UsersList
//		Список читателей
// Security:
// - basic
// Responses:
// 	 200: body:User Список читателей успешно получен

// swagger:parameters UsersList

// swagger:model User
type UsersList struct {
	// in:body
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//go:generate `swagger generate spec -o ./public/swagger.json --scan-models`
