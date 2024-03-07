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
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	BirthDate string        `json:"birthDate"`
	Books     []*model.Book `json:"books"`
}

// swagger:route GET /library/popular-authors author AuthorsTop
//		Список популярных авторов
// security:
// - basic
// responses:
// 	 200: body:Author Cписок популярных авторов успешно получен

// swagger:model
type Authors []model.Author

// swagger:route POST /library/author author AuthorAdd
//		Добавить автора
// security:
// - basic
// responses:
// 	 200: body:Author Автор успешно добавлен

// swagger:model Author
type AuthorAdd struct{}
