package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/x6txy/go2024/finalproject/database/postgres"
	"github.com/x6txy/go2024/finalproject/model"
	"github.com/x6txy/go2024/finalproject/service"
)

var dbPool *sql.DB

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userService := getServiceFromContext(r)
	users, err := userService.GetAllUsers()
	if err != nil {
		log.Printf("Error retrieving users: %v", err)
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("id")
	if userIDParam == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userService := getServiceFromContext(r)
	user, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userService := getServiceFromContext(r)
	userID, err := userService.CreateUser(&newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
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

func getServiceFromContext(r *http.Request) *service.UserService {
	if dbPool == nil {
		dbPool = postgres.InitDB()
	}

	return service.NewUserService(dbPool)
}
