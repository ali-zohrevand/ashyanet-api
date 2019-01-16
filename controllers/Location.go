package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"net/http"
)

func CreateLocation(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	location := new(models.Location)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.CreateLocation(location)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
