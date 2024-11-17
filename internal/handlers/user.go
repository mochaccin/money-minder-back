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
	cardRepository = &repositories.CardRepo{
		MongoCollection: service.GetCollection("cards"),
	}
	spendRepository = &repositories.SpendRepo{
		MongoCollection: service.GetCollection("spends"),
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

func GetUserByID(w http.ResponseWriter, r *http.Request) error {

	id := r.PathValue("id")

	user, err := userRepository.FindUserByID(id)

	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, user)
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) error {

	userId := r.PathValue("id")

	updateUsr := &UpdateUserPasswordRequest{}
	derr := json.NewDecoder(r.Body).Decode(updateUsr)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt update user password, verify that the values are formatted correctly",
		}
	}

	err := userRepository.UpdatePassword(userId, updateUsr.NewPassword)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "User's password updated succesfully")
}

func UpdateUserName(w http.ResponseWriter, r *http.Request) error {

	userId := r.PathValue("id")

	updateUsr := &UpdateUserNameRequest{}
	derr := json.NewDecoder(r.Body).Decode(updateUsr)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt update user username, verify that the values are formatted correctly",
		}
	}

	err := userRepository.UpdateName(userId, updateUsr.NewUsername)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "User's username updated succesfully")
}

func AddUserCard(w http.ResponseWriter, r *http.Request) error {

	userId := r.PathValue("id")

	addCardRequest := &UserCardRequest{}
	derr := json.NewDecoder(r.Body).Decode(addCardRequest)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt add card to user, verify that the values are formatted correctly",
		}
	}

	err := userRepository.AddCard(userId, addCardRequest.CardId, cardRepository)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "New card added sucessfully.")
}

func RemoveUserCard(w http.ResponseWriter, r *http.Request) error {

	userId := r.PathValue("id")

	removeCardRequest := &UserCardRequest{}
	derr := json.NewDecoder(r.Body).Decode(removeCardRequest)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt remove card to from, verify that the values are formatted correctly",
		}
	}

	err := userRepository.RemoveCard(userId, removeCardRequest.CardId, cardRepository)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "card deleted sucessfully.")
}

func AddUserSpend(w http.ResponseWriter, r *http.Request) error {

	userId := r.PathValue("id")

	addSpendRequest := &UserSpendRequest{}
	derr := json.NewDecoder(r.Body).Decode(addSpendRequest)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt add spend to user, verify that the values are formatted correctly",
		}
	}

	err := userRepository.AddSpend(userId, addSpendRequest.SpendId, spendRepository)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "New spend added sucessfully.")
}

func RemoveUserSpend(w http.ResponseWriter, r *http.Request) error {

	userId := r.PathValue("id")

	removeSpendRequest := &UserSpendRequest{}
	derr := json.NewDecoder(r.Body).Decode(removeSpendRequest)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Couldnt remove spend to from, verify that the values are formatted correctly",
		}
	}

	err := userRepository.RemoveSpend(userId, removeSpendRequest.SpendId, spendRepository)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, "spend deleted sucessfully.")
}

type UpdateUserPasswordRequest struct {
	NewPassword string `json:"newPassword" bson:"new_password"`
}

type UpdateUserNameRequest struct {
	NewUsername string `json:"newUsername" bson:"new_username"`
}

type UserCardRequest struct {
	CardId string `json:"cardId" bson:"card_id"`
}

type UserSpendRequest struct {
	SpendId string `json:"spendId" bson:"spend_id"`
}
