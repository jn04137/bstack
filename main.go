package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"com/bstack/controller"
	"com/bstack/dependencies"
)

func main() {
	env := dependencies.CreateEnvironment()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	userController := controller.NewUserController(env)
	r.Mount("/user", userController.GetRoute())

	http.ListenAndServe(":8080", r)
}
