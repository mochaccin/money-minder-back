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
