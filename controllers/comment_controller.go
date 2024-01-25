package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment

	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentID, err := services.CreateComment(newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error while creating the comment."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("Comment created correctly with ID: %d\n", commentID)
	w.Write([]byte(response))
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	movieID, _ := strconv.Atoi(r.URL.Query().Get("movie_id"))
	content := r.URL.Query().Get("content")

	comments, err := services.GetComments(userID, movieID, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error fetching the comments."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func GetCommentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	comment, err := services.GetCommentByID(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var updatedComment models.Comment
	err = json.NewDecoder(r.Body).Decode(&updatedComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateComment(commentID, updatedComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Comment with ID %d updated successfully.", commentID)
	w.Write([]byte(response))
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteComment(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Comment with ID %d deleted successfully.", commentID)
	w.Write([]byte(response))
}
