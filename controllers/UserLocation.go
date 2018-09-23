package controllers

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
	"net/http"
)

func AddUserToLocation(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	device := new(models.UserLocation)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&device)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.AddUserLocation(device)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
