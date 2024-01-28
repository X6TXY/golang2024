package main

import (
    "log"
    "net/http"
    "github.com/x6txy/tsis1/internal/handlers"
    "github.com/gorilla/mux"
)

func main() {
    log.Println("Starting API server")
    router := mux.NewRouter()

    log.Println("Creating routes")
    router.HandleFunc("/players", handlers.GetPlayers).Methods("GET")
    router.HandleFunc("/players/last_name/{last_name}", handlers.GetPlayerByName).Methods("GET")
	router.HandleFunc("/players/{id}", handlers.GetPlayerById).Methods("GET")
    router.HandleFunc("/health-checking", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/players/position/{position}", handlers.GetPlayersByPosition).Methods("GET")
	router.HandleFunc("/players/nation/{nation}", handlers.GetPlayersByNation).Methods("GET")
    http.Handle("/", router)

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
