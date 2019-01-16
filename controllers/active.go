package controllers

import (
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
	"net/http"
)

func Active(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName:= vars["user"]
	activeCode:= vars["activeCode"]
	w.Header().Set("Content-Type", "application/json")
	responseStatus, token := services.Active(userName,activeCode)
	w.WriteHeader(responseStatus)
	w.Write(token)
}
