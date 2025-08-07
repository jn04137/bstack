package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"

	"com/bstack/dependencies"
	"com/bstack/models"
	"com/bstack/repositories"
	"com/bstack/services"
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

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.NanoId = nanoId

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
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
	passwordInput := user.Password

	user, err = userRepo.GetUser(user.Username)
	if err != nil {
		log.Printf("Failed to get user with error: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInput))
	if err != nil {
		log.Printf("Credentials failed: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ss, err := services.CreateJwtToken(user)
	if err != nil {
		log.Printf("Failed making jwt token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userCookie := services.CreateJwtCookie(ss)
	http.SetCookie(w, &userCookie)

	log.Printf("This user logged in: %s", user.Username)
	w.WriteHeader(http.StatusOK)
	return
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


