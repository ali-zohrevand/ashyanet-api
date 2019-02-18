package controllers

import (
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/gorilla/mux"
	"net/http"
)

func Jwt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jwt := vars["jwt"]
	responseStatus, token := services.IsJwtValid(jwt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}
