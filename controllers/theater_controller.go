package controllers

import (
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
	"net/http"
	"strconv"
)

func CreateTheater(w http.ResponseWriter, r *http.Request) {
	var newTheater models.Theater
	CreateControllerItem(w, r, &newTheater, func(data interface{}) (int, error) {
		theater := data.(*models.Theater)
		return services.CreateTheater(*theater)
	}, "Theater")
}

func GetTheaters(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	capacity, _ := strconv.Atoi(r.URL.Query().Get("capacity"))
	capacityGt, _ := strconv.Atoi(r.URL.Query().Get("capacity_gt"))
	capacityLt, _ := strconv.Atoi(r.URL.Query().Get("capacity_lt"))

	GetControllerItems(w, r, func() (interface{}, error) {
		return services.GetTheaters(name, capacity, capacityGt, capacityLt)
	}, "Theaters")
}

func GetTheaterByID(w http.ResponseWriter, r *http.Request) {
	GetControllerItemByID(w, r, func(id int) (interface{}, error) {
		return services.GetTheaterByID(id)
	}, "Theater")
}

func UpdateTheater(w http.ResponseWriter, r *http.Request) {
	var updatedTheater models.Theater
	UpdateControllerItem(w, r, &updatedTheater, func(id int, data interface{}) error {
		theater := data.(*models.Theater)
		return services.UpdateTheater(id, *theater)
	}, "Theater")
}

func DeleteTheater(w http.ResponseWriter, r *http.Request) {
	DeleteControllerItem(w, r, func(id int) error {
		return services.DeleteTheater(id)
	}, "Theater")
}
