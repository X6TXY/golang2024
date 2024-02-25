//Authentication handler for authentication

package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/x6txy/go2024/finalproject/database/postgres"
	"github.com/x6txy/go2024/finalproject/model"
	"github.com/x6txy/go2024/finalproject/service"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Printf("Error decoding registration request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	newUser.Password = hashedPassword

	userService := getUserServiceFromContext(r)
	userID, err := userService.RegisterUser(&newUser)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	response := map[string]int{"id": userID}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Printf("Error decoding login request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userService := getUserServiceFromContext(r)
	user, err := userService.GetUserByUsername(loginRequest.Username)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(loginRequest.Password))
	if err != nil {
		log.Printf("Error comparing passwords: %v", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authentication successful"))
}
func getUserServiceFromContext(r *http.Request) *service.UserService {
	if dbPool == nil {
		dbPool = postgres.InitDB()
	}

	return service.NewUserService(dbPool)
}

func AllUsersHandler(w http.ResponseWriter, r *http.Request) {
	userService := getUserServiceFromContext(r)

	users, err := userService.GetAllUsers()
	if err != nil {
		log.Printf("Error retrieving all users: %v", err)
		http.Error(w, "Failed to retrieve all users", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(users)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
