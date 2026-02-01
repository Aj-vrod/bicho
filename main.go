package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

}

func startServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", rootHandler)

	http.ListenAndServe(":3000", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
