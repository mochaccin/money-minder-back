package handlers

import (
	"encoding/json"
	"money-minder/internal/database"
	"money-minder/internal/repositories"
	"money-minder/internal/types"
	"net/http"
)

var (
	service        = database.New()
	userRepository = &repositories.UserRepo{
		MongoCollection: service.GetCollection("users"),
	}
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

	result, err := userRepository.InsertUser(usr)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, result)
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
