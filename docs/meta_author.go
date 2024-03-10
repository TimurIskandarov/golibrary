package docs

import "golibrary/internal/model"

// swagger:route GET /library/authors author AuthorsList
//		Список авторов
// responses:
// 	 200: body:[]Author Cписок авторов успешно получен


// swagger:route GET /library/popular-authors author AuthorsTop
//		Список популярных авторов
// responses:
// 	 200: body:[]Author Cписок популярных авторов успешно получен

// swagger:route POST /library/author author AuthorAdd
//		Добавить автора
// responses:
// 	 200: body:Author Автор успешно добавлен

// swagger:model
type Author struct {        
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	BirthDate string        `json:"birthDate"`
	Books     []*model.Book `json:"books"`
}
