package docs

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