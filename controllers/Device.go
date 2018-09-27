package controllers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
)

func CreateDevice(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	device := new(models.DeviceInDB)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&device)
	/*	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}*/
	w.Header().Set("Content-Type", "application/json")
	//................................
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	responseStatus, token := services.CreateDevice(device, userInDB)
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
