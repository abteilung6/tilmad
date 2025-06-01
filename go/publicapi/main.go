package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/abteilung6/tilmad/go/publicapi/endpoints"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Get("/blackberries", endpoints.GetBlackberries)

	http.ListenAndServe(":8080", r)
}
