package handlers

import (
	"encoding/json"
	"money-minder/internal/types"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	usr := &types.User{}
	derr := json.NewDecoder(r.Body).Decode(usr)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt create user, verify that the values are formatted corrected",
		}
	}
	return writeJSON(w, http.StatusOK, usr)
}

type APIError struct {
	Status int
	Msg    string
}

func (e APIError) Error() string {
	return e.Msg
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
