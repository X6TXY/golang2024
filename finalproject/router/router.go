package router

import (
	"github.com/gorilla/mux"
	"github.com/x6txy/go2024/finalproject/api/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/comments", handler.GetComments).Methods("GET")
	r.HandleFunc("/comments/{id}", handler.GetComment).Methods("GET")
	r.HandleFunc("/comments", handler.CreateComment).Methods("POST")
	r.HandleFunc("/comments/{id}", handler.UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/{id}", handler.DeleteComment).Methods("DELETE")

	return r
}
