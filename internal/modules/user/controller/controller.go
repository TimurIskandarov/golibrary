package controller

import (
	"encoding/json"
	"net/http"

	"golibrary/internal/modules/user/repository"
	"golibrary/internal/modules/user/service"
	"golibrary/internal/model"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Userer interface {
	Add(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type User struct {
	service service.Userer
}

func NewUser(db *sqlx.DB, logger *zap.Logger) Userer {
	return &User{
		service: service.NewUserService(
			repo.NewUserRepository(db),
			logger,
		),
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
