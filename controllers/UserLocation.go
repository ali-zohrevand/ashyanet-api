package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
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
