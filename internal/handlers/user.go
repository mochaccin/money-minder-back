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
			Msg:    "Couldnt create user, verify that the values are formatted correctly",
		}
	}

	result, err := userRepository.InsertUser(usr)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, result)
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) error {
	updateUsr := &UpdateUserPasswordRequest{}
	derr := json.NewDecoder(r.Body).Decode(updateUsr)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt update user password, verify that the values are formatted correctly",
		}
	}

	err := userRepository.UpdatePassword(updateUsr.userId, updateUsr.newPassword)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "User's password updated succesfully")
}

func UpdateUserName(w http.ResponseWriter, r *http.Request) error {
	updateUsr := &UpdateUserNameRequest{}
	derr := json.NewDecoder(r.Body).Decode(updateUsr)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt update user username, verify that the values are formatted correctly",
		}
	}

	err := userRepository.UpdateName(updateUsr.userId, updateUsr.newUsername)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "User's username updated succesfully")
}

type UpdateUserPasswordRequest struct {
	userId      string
	newPassword string
}

type UpdateUserNameRequest struct {
	userId      string
	newUsername string
}
