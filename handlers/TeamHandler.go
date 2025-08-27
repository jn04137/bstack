package handlers

import (
	"com/bstack/dependencies"
	"com/bstack/repositories"
	"encoding/json"
	"log"
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

func (handler TeamHandler) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint was hit")
	db := handler.teamRepo
	teams, err := db.GetAllTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teams)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
