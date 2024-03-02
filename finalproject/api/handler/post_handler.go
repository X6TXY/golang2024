package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/x6txy/go2024/finalproject/database/postgres"
	"github.com/x6txy/go2024/finalproject/model"
	"github.com/x6txy/go2024/finalproject/service"
)

var dbPool *sql.DB

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost model.Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newPost.Date = time.Now().Truncate(time.Minute)
	newPost.Update_date = time.Now().Truncate(time.Minute)
	postService := getServiceFromContext(r)

	postID, err := postService.CreatePost(&newPost)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	response := map[string]int{"id": postID}
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

func GetPosts(w http.ResponseWriter, r *http.Request) {
	postService := getServiceFromContext(r)
	posts, err := postService.GetAllPosts()
	if err != nil {
		log.Printf("Error retrieving posts: %v", err)
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(posts)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIDStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing post ID in request", http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	postService := getServiceFromContext(r)

	post, err := postService.GetPostByID(postID)
	if err != nil {
		log.Printf("Error retrieving post by ID: %v", err)
		http.Error(w, "Failed to retrieve post by ID", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(post)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func DeletePostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIDStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing post ID in request", http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	postService := getServiceFromContext(r)
	err = postService.DeletePostById(postID)
	if err != nil {
		log.Printf("Error deleting post: %v", err)
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIDStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing post ID in request", http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var updateRequest struct {
		Text string `json:"text"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}

	postService := getServiceFromContext(r)

	err = postService.UpdatePost(postID, updateRequest.Text)
	if err != nil {
		log.Printf("Error updating post: %v", err)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getServiceFromContext(r *http.Request) *service.PostService {
	if dbPool == nil {
		dbPool = postgres.InitDB()
	}

	return service.NewPostService(dbPool)
}
