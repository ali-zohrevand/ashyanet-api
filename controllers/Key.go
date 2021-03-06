package controllers

import (
	"github.com/ali-zohrevand/ashyanet-api/services"
	"net/http"
)

func CreatKey(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.CreatValidKey()
	w.WriteHeader(responseStatus)
	w.Write(token)
}
