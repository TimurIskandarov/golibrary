package docs

import "golibrary/internal/model"

// swagger:route GET /library/authors author Author
//		Список авторов
// security:
// - basic
// responses:
// 	 200: body:Author Cписок авторов успешно получен

// swagger:parameters Author

// swagger:model Author
type Author struct {
	// in:body
	ID        int           `json:"id"`
	// in:body
	Name      string        `json:"name"`
	// in:body
	BirthDate string        `json:"birthDate"`
	// in:body
	Books     []*model.Book `json:"books"`
}

// swagger:route GET /library/popular-authors author AuthorsTop
//		Список популярных авторов
// security:
// - basic
// responses:
// 	 200: body:Author Cписок популярных авторов успешно получен

// swagger:parameters AuthorsTop

// swagger:model Authors
type Authors []model.Author

// swagger:route POST /library/author author AuthorAdd
//		Добавить автора
// security:
// - basic
// responses:
// 	 200: body:Author Автор успешно добавлен

// swagger:parameters AuthorAdd
type _ struct {
	//in:body
	Author Author
}

// swagger:model
type AuthorAdd struct {}
