package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Blackberries struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetBlackberries(w http.ResponseWriter, r *http.Request) {
	blackberries := []Blackberries{
		{ID: 1, Name: "Triple Crown"},
		{ID: 2, Name: "Black Satin"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blackberries)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Get("/blackberries", GetBlackberries)

	http.ListenAndServe(":8080", r)
}
