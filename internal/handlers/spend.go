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
