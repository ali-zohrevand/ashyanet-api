package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/gorilla/mux"
)

func User(router *mux.Router) *mux.Router {
	router.HandleFunc("/register", controllers.RegisterPost).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	return router
}
