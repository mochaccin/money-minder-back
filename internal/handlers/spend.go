package handlers

import (
	"encoding/json"
	"money-minder/internal/types"
	"net/http"
)

func CreateSpend(w http.ResponseWriter, r *http.Request) error {
	spend := &types.Spend{}
	derr := json.NewDecoder(r.Body).Decode(spend)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Could not create spend, verify that the values are formatted correctly",
		}
	}

	result, err := spendRepository.InsertSpend(spend)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, result)
}

func DeleteSpend(w http.ResponseWriter, r *http.Request) error {

	spendId := r.PathValue("id")

	result, err := spendRepository.DeleteSpend(spendId)

	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, result)
}

func GetSpendByID(w http.ResponseWriter, r *http.Request) error {

	spendId := r.PathValue("id")

	spend, err := spendRepository.FindSpendByID(spendId)

	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, spend)
}

func GetAllSpendsByUserID(w http.ResponseWriter, r *http.Request) error {

	userId := r.PathValue("id")

	cards, err := spendRepository.GetSpendsByUserID(userId)

	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, cards)
}

func GetAllSpendsByCardID(w http.ResponseWriter, r *http.Request) error {

	cardId := r.PathValue("id")

	cards, err := spendRepository.GetSpendsByCardID(cardId)

	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, cards)
}
