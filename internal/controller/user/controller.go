package controller

import (
	"encoding/json"
	"net/http"

	"golibrary/internal/model"
	"golibrary/internal/service/user"
)

type Userer interface {
	Add(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type User struct {
	service service.Userer
}

func NewUser(service service.Userer) Userer {
	return &User{
		service: service,
	}
}

// List implements Userer.
func (u *User) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Add implements Userer.
func (u *User) Add(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID, err = u.service.Add(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
