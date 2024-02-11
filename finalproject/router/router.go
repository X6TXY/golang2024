package router

import (
	"github.com/gorilla/mux"
	"github.com/x6txy/go2024/finalproject/api/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/api/users", handler.CreateUser).Methods("POST")

	return r
}