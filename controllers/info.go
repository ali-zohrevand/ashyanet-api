package controllers

import (
	"github.com/ali-zohrevand/ashyanet-api/services"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.Info(userInDB)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}

}
