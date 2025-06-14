package handlers

import (
	"log"
	"net/http"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"github.com/matoous/go-nanoid/v2"

	"com/bstack/dependencies"
	"com/bstack/repositories"
	"com/bstack/models"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(env *dependencies.Environment) *UserHandler {
	return &UserHandler{
		userRepo: repositories.NewUserRepository(env),
	}
}

func (handler UserHandler) UserSignup(w http.ResponseWriter, r *http.Request){
	userRepo := handler.userRepo
	user := models.UserAccount{}

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)
	if err != nil {
		log.Printf("Failed to parse body: %v", err)
	}
	nanoId, err := gonanoid.New()
	if err != nil {
		log.Printf("Error generating nanoId: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.NanoId = nanoId

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		log.Printf("Error occurred while hashing password: %v", err)
	}
	user.Password = string(hash)
	
	err = userRepo.CreateUser(user)
	if err != nil {
		log.Printf("Error trying to create user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler UserHandler) UserSignin(w http.ResponseWriter, r *http.Request) {
	userRepo := handler.userRepo
	user := models.UserAccount{}

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)

	if err != nil {
		log.Printf("Failed to decode request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err = userRepo.GetUser(user.Username)
	if err != nil {
		log.Printf("Failed to get user with error: %v", err)
	}
	log.Printf("This user logged in: %s", user.Username)
}

func (handler UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userRepo := handler.userRepo

	users, err := userRepo.GetAllUsers()
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}


