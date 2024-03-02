package router

import (
	"github.com/gorilla/mux"
	"github.com/x6txy/go2024/finalproject/api/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/posts", handler.CreatePost).Methods("POST")
	r.HandleFunc("/posts", handler.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", handler.GetPostById).Methods("GET")
	r.HandleFunc("/posts/{id}", handler.DeletePostById).Methods("DELETE")
	r.HandleFunc("/posts/{id}", handler.UpdatePostHandler).Methods("PUT")

	
	return r
}
