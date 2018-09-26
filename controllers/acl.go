package controllers

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
	"net/http"
)

func CreateAcl(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	Acl := new(models.MqttAcl)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&Acl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.EmqttHttpCreateAcl(Acl)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
