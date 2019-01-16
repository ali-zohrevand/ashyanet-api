package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"net/http"
)

func AddUserToDevice(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	device := new(models.UserDevice)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&device)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.AddUserDevice(device)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
