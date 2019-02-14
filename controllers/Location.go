package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/gorilla/mux"
	"net/http"
)

func LocationCreate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	location := new(models.Location)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user not found"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.CreateLocation(location, userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}

}

func LocationGetAll(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user not found"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.LocationGetAllByUsername(userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}

}
func LocationDeleteById(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	id := vars["id"]

	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user not found"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.LocationDeleteById(id, userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
