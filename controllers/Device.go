package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
)

func CreateDevice(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	device := new(models.Device)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&device)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.CreateDevice(device)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
func AddUserToDevice(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	device := new(models.UserDevice)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&device)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.AddUserToDevice(device)
	w.WriteHeader(responseStatus)
	w.Write(token)
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
