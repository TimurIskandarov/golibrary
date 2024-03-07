package docs

// swagger:parameters AuthorsTop
// swagger:parameters AuthorAdd
type _ struct {
	//in:body
	Author Author
}

// swagger:parameters BooksList
// swagger:parameters BookTake
type _ struct {
	//in:path
	BookId int `json:"bookId"`
	//in:path
	UserId int `json:"userId"`
}
// swagger:parameters BookReturn
type _ struct {
	//in:path
	BookId int `json:"bookId"`
	//in:path
	UserId int `json:"userId"`
}
// swagger:parameters BookAdd
type _ struct {
	// in:body
	Book Book
}

// swagger:parameters UsersList
// swagger:parameters UserAdd
type _ struct {
	// in:body
	User User
}
