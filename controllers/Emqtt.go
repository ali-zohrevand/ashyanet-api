package controllers

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
	"net/http"
)

func CreateEmqttUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(models.MqttUser)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.CreateEmqttUser(requestUser)
	w.WriteHeader(responseStatus)
	w.Write(token)

}
