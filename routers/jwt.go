package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/gorilla/mux"
)

func Jwt(router *mux.Router) *mux.Router {
	router.HandleFunc("/jwt/{jwt}", controllers.Jwt).Methods("GET")
	return router
}
