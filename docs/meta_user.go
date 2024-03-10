package docs

// swagger:route GET /library/users user UsersList
//		Список читателей
// responses:
// 	 200: body:[]User Список читателей успешно получен

// swagger:route POST /library/user user UserAdd
//		Добавить читателя
// responses:
// 	 200: body:User Читатель успешно добавлен

// swagger:model
type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Books []*Book `json:"books,omitempty"`
}
