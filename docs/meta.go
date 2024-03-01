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

//go:generate `swagger generate spec -o ./public/swagger.json --scan-models`

// swagger:route POST /api/register auth RegisterRequest
//		Регистрация нового пользователя
// Security:
// - basic
// Responses:
// 	 200: body:RegisterResponse

// swagger:parameters RegisterRequest
type RegisterRequest struct {
	// Регистрация
	// in:body
	// required:true
	// example: {"name":"tim","password":"123"}
	Registration string
}

// swagger:model RegisterResponse
type RegisterResponse struct {
	// in:body
	// Token содержит информацию о регистрации
	Token string
}

// swagger:route POST /api/login auth LoginRequest
//		Авторизация пользователя
// Security:
// - basic
// Responses:
// 	 200: body:LoginResponse

// swagger:parameters LoginRequest
type LoginRequest struct {
	// Авторизация
	// in:body
	// required:true
	// example: {"name":"tim","password":"123"}
	Authorization string
}

// swagger:model LoginResponse
type LoginResponse struct {
	// in:body
	// Token содержит информацию о токене
	Token string
}
