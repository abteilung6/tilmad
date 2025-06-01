package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/abteilung6/tilmad/go/publicapi/endpoints"
)

type Server struct {
	Router *chi.Mux
}

func main() {
	server := CreateNewServer()
	MountHandlers(server)
	http.ListenAndServe(":8080", server.Router)
}

func CreateNewServer() *Server {
	server := &Server{}
	server.Router = chi.NewRouter()
	return server
}

func MountHandlers(server *Server) {
	server.Router.Use(middleware.Logger)
	server.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	server.Router.Get("/blackberries", endpoints.GetBlackberries)
}
