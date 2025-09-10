package handlers

import (
	"com/bstack/dependencies"
	"com/bstack/repositories"
	"com/bstack/models"
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
	var team models.Team
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&team)

	err := db.CreateTeam(team.TeamName, team.OwnerNanoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (handler TeamHandler) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	db := handler.teamRepo
	teams, err := db.GetAllTeams()
	if err != nil {
		log.Printf("Error getting teams: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teams)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (handler TeamHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	//var team models.Team

	//err := decoder.Decode(&team)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	//db := handler.teamRepo
}

func (handler TeamHandler) EditTeam(w http.ResponseWriter, r *http.Request) {

}

func (handler TeamHandler) JoinTeam(w http.ResponseWriter, r *http.Request) {

}

func (handler TeamHandler) AcceptJoinTeamRequest(w http.ResponseWriter, r *http.Request) {

}
