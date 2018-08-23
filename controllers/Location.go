package controllers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
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
