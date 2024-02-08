package controllers

import (
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
	"net/http"
	"strconv"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment
	CreateControllerItem(w, r, &newComment, func(data interface{}) (int, error) {
		comment := data.(*models.Comment)
		return services.CreateComment(*comment)
	}, "Comment")
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	movieID, _ := strconv.Atoi(r.URL.Query().Get("movie_id"))
	content := r.URL.Query().Get("content")

	GetControllerItems(w, r, func() (interface{}, error) {
		return services.GetComments(userID, movieID, content)
	}, "Comments")
}

func GetCommentByID(w http.ResponseWriter, r *http.Request) {
	GetControllerItemByID(w, r, func(id int) (interface{}, error) {
		return services.GetCommentByID(id)
	}, "Comment")
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var updatedComment models.Comment
	UpdateControllerItem(w, r, &updatedComment, func(id int, data interface{}) error {
		comment := data.(*models.Comment)
		return services.UpdateComment(id, *comment)
	}, "Comment")
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	DeleteControllerItem(w, r, func(id int) error {
		return services.DeleteComment(id)
	}, "Comment")
}
