package handlers

import (
	"encoding/json"
	"money-minder/internal/types"
	"net/http"
)

func CreateCard(w http.ResponseWriter, r *http.Request) error {
	card := &types.Card{}
	derr := json.NewDecoder(r.Body).Decode(card)

	if derr != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "Could not create card, verify that the values are formatted correctly",
		}
	}

	result, err := cardRepository.InsertCard(card)
	if err != nil {
		return APIError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	return WriteJSON(w, http.StatusOK, result)
}
