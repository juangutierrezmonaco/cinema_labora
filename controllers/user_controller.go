package controllers

import (
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	CreateControllerItem(w, r, &newUser, func(data interface{}) (int, error) {
		user := data.(*models.User)
		return services.CreateUser(*user)
	}, "User")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	email := r.URL.Query().Get("email")
	gender := r.URL.Query().Get("gender")

	GetControllerItems(w, r, func() (interface{}, error) {
		return services.GetUsers(firstName, lastName, email, gender)
	}, "Users")
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	GetControllerItemByID(w, r, func(id int) (interface{}, error) {
		return services.GetUserByID(id)
	}, "User")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	UpdateControllerItem(w, r, &updatedUser, func(id int, data interface{}) error {
		user := data.(*models.User)
		return services.UpdateUser(id, *user)
	}, "User")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	DeleteControllerItem(w, r, func(id int) error {
		return services.DeleteUser(id)
	}, "User")
}
