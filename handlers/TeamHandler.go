package handlers

import (
	"com/bstack/dependencies"
	"com/bstack/repositories"
	"net/http"
)

type TeamHandler struct {
	teamRepo *repositories.TeamRepository
}

func NewTeamHandler(env *dependencies.Environment) *TeamHandler {
	return &TeamHandler{
		teamRepo: repositories.NewTeamRepository(env),
	}
}

func (handler TeamHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	db := handler.teamRepo

	err := db.CreateTeam("teamName", "ownerNanoId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
