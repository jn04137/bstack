package controller

import (
	"com/bstack/dependencies"
	"com/bstack/handlers"

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

func (controller TeamController) GetRoute() *chi.Mux {
	r := chi.NewRouter()
	handler := controller.teamHandler
	r.Get("/allTeams", handler.GetAllTeams)
	r.Get("/{teamNanoId}", handler.TeamPageData)

	protectedRoutes := chi.NewRouter()
	protectedRoutes.Get("/createTeam", handler.CreateTeam)

	r.Mount("/protected", protectedRoutes)
	return r
}
