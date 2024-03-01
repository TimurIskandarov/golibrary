package responder

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})
}

type Respond struct {
	log *zap.Logger
}

func NewResponder(logger *zap.Logger) Responder {
	return &Respond{log: logger}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		r.log.Error("responder json encode error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
