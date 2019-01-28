package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/gorilla/mux"
)

func Register(router *mux.Router) *mux.Router {
	router.HandleFunc("/register", controllers.RegisterPost).Methods("POST")

	return router
}
