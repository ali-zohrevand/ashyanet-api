package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/gorilla/mux"
)

func login(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	return router
}
