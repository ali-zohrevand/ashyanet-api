package controllers

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
	"net/http"
)

func EventCreate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	dataBinde := new(models.DataBind)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dataBinde)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	responseStatus, token := services.EventCreate(*dataBinde, userInDB)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
