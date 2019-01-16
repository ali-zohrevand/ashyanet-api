package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"net/http"
)

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.Register(requestUser)
	w.WriteHeader(responseStatus)
	w.Write(token)

}
