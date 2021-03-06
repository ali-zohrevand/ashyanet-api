package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/gorilla/mux"
	"net/http"
)

func TypesCreate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	typeObj := new(models.Types)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&typeObj)
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user not found"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.TypeCreate(*typeObj, userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}

}
func TypesDelete(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	id := vars["id"]

	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user not found"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.TypesDeleteById(id, userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}

}

func TypesGetAll(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user not found"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.TypeGetAllByUsername(userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}

}
