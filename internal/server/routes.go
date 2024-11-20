package server

import (
	"encoding/json"
	"log"
	"log/slog"
	"money-minder/internal/handlers"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.healthHandler)
	mux.HandleFunc("POST /users", makeHandler(handlers.CreateUser))
	mux.HandleFunc("POST /cards", makeHandler(handlers.CreateCard))
	mux.HandleFunc("DELETE /cards/{id}", makeHandler(handlers.DeleteCard))
	mux.HandleFunc("GET /cards/{id}", makeHandler(handlers.GetCardByID))
	mux.HandleFunc("GET /users/{id}/cards", makeHandler(handlers.GetAllCardsByUserID))
	mux.HandleFunc("POST /spends", makeHandler(handlers.CreateSpend))
	mux.HandleFunc("GET /users/{id}", makeHandler(handlers.GetUserByID))
	mux.HandleFunc("PUT /users/{id}/password", makeHandler(handlers.UpdateUserPassword))
	mux.HandleFunc("PUT /users/{id}/username", makeHandler(handlers.UpdateUserName))
	mux.HandleFunc("POST /users/{id}/cards", makeHandler(handlers.AddUserCard))
	mux.HandleFunc("DELETE /users/{id}/cards", makeHandler(handlers.RemoveUserCard))
	mux.HandleFunc("POST /users/{id}/spends", makeHandler(handlers.AddUserSpend))
	mux.HandleFunc("DELETE /users/{id}/spends", makeHandler(handlers.RemoveUserSpend))
	mux.HandleFunc("POST /cards/{id}/spends", makeHandler(handlers.AddCardSpend))
	mux.HandleFunc("DELETE /cards/{id}/spends", makeHandler(handlers.RemoveCardSpend))
	mux.HandleFunc("DELETE /spends/{id}", makeHandler(handlers.DeleteSpend))
	mux.HandleFunc("GET /spends/{id}", makeHandler(handlers.GetSpendByID))
	mux.HandleFunc("GET /users/{id}/spends", makeHandler(handlers.GetAllSpendsByUserID))
	mux.HandleFunc("GET /cards/{id}/spends", makeHandler(handlers.GetAllSpendsByCardID))

	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHandler(h apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if e, ok := err.(handlers.APIError); ok {
				slog.Error("API Error", "err", e, "status", e.Status)
				handlers.WriteJSON(w, e.Status, e)
			}
		}
	}
}
