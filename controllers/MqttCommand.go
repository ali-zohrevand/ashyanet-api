package controllers

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
	"net/http"
)

func MqttCommand(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	command := new(models.Command)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	responseStatus, token := services.MqttHttpCommand(*command, userInDB)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
