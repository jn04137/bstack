package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"

	"com/bstack/controller"
	"com/bstack/dependencies"
)

func main() {
	env := dependencies.CreateEnvironment()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	userController := controller.NewUserController(env)
	teamController := controller.NewTeamController(env)
	r.Mount("/user", userController.GetRoute())
	r.Mount("/team", teamController.GetRoute())

	http.ListenAndServe(":8080", r)
}
