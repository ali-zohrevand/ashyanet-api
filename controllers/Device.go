package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
)

func DeviceCreate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	device := new(models.Device)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&device)
	/*	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}*/
	//................................
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user not found"))

	} else {
		w.Header().Set("Content-Type", "application/json")

		responseStatus, token := services.DeviceCreate(device, userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}

}
func DeviceGetAll(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.DevicesGetAll(userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
func DeviceGetId(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	id := vars["id"]
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.DeviceGetId(userInDB, id)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
func DeviceDeleteId(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	id := vars["id"]
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.DeviceDeleteID(userInDB, id)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
func DeviceUpdateId(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	device := new(models.Device)
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&device)
	fmt.Print(errDecode)
	userInDB, err := GetUserFromHeader(r)
	vars := mux.Vars(r)
	id := vars["id"]
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {

		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.DeviceUpdateID(id, userInDB, *device)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
func AddKeyToDevice(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	deviceKey := new(models.DeviceKey)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deviceKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.AddKeyToDevice(deviceKey)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
