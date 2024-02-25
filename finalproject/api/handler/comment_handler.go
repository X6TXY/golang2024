//Handler for CRUD of comments

package handler

import (
	"database/sql"
	"encoding/json"
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

func GetComments(w http.ResponseWriter, r *http.Request) {
	commentService := getServiceFromContext(r)
	comments, err := commentService.GetAllComments()
	if err != nil {
		log.Printf("Error retrieving comments: %v", err)
		http.Error(w, "Failed to retrieve comments", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(comments)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	commentIDParam := r.URL.Query().Get("id")
	if commentIDParam == "" {
		http.Error(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	commentID, err := strconv.Atoi(commentIDParam)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	commentService := getServiceFromContext(r)
	comment, err := commentService.GetCommentByID(commentID)
	if err != nil {
		log.Printf("Error retrieving comment: %v", err)
		http.Error(w, "Failed to retrieve comment", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(comment)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var newComment model.Comment
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newComment.Date = time.Now().Truncate(time.Minute)

	commentService := getServiceFromContext(r)
	commentID, err := commentService.CreateComment(&newComment)
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	response := map[string]int{"id": commentID}
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

func getServiceFromContext(r *http.Request) *service.CommentService {
	if dbPool == nil {
		dbPool = postgres.InitDB()
	}

	return service.NewCommentService(dbPool)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentIDParam := mux.Vars(r)["id"]
	if commentIDParam == "" {
		http.Error(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	commentID, err := strconv.Atoi(commentIDParam)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	commentService := getServiceFromContext(r)
	err = commentService.DeleteCommentByID(commentID)
	if err != nil {
		log.Printf("Error deleting comment: %v", err)
		http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	commentIDParam := mux.Vars(r)["id"]
	if commentIDParam == "" {
		http.Error(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	commentID, err := strconv.Atoi(commentIDParam)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var updatedComment model.Comment
	err = json.NewDecoder(r.Body).Decode(&updatedComment)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	commentService := getServiceFromContext(r)

	existingComment, err := commentService.GetCommentByID(commentID)
	if err != nil {
		log.Printf("Error retrieving comment: %v", err)
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}


	now := time.Now()
	updatedComment.Date = time.Date(now.Year(), now.Month(), now.Day(), existingComment.Date.Hour(), existingComment.Date.Minute(), 0, 0, now.Location())

	updatedComment.ID = commentID
	err = commentService.UpdateComment(&updatedComment)
	if err != nil {
		log.Printf("Error updating comment: %v", err)
		http.Error(w, "Failed to update comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}