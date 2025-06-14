package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"com/bstack/dependencies"
	"com/bstack/handlers"
)

type UserController struct {
	userHandler *handlers.UserHandler
}

func NewUserController(env *dependencies.Environment) *UserController {
	return &UserController{
		userHandler: handlers.NewUserHandler(env),  
	}
}

func (controller UserController) GetRoute() *chi.Mux {
	r := chi.NewRouter()
	userHandler := controller.userHandler

	r.Get("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the user endpoint"))
	})
	
	r.Post("/signup", userHandler.UserSignup)
	r.Post("/login", userHandler.UserSignin)
	r.Get("/getAll", userHandler.GetAllUsers)

	return r
}
