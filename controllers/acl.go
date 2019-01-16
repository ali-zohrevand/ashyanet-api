package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
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
