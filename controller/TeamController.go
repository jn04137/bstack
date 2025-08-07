package controller

import (
	"com/bstack/dependencies"
	"com/bstack/handlers"
	"com/bstack/middlewares"

	"github.com/go-chi/chi/v5"
)

type TeamController struct {
	teamHandler *handlers.TeamHandler
}

func NewTeamController(env *dependencies.Environment) *TeamController {
	return &TeamController{
		teamHandler: handlers.NewTeamHandler(env),
	}
}

func (controller TeamController) GetRoute() {
	r := chi.NewRouter()
	r.Use(middlewares.UserAuthMiddleware)

	handler := controller.teamHandler

	r.Post("/createTeam", handler.CreateTeam)
}
