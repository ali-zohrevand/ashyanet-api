package routers

import (
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/controllers"
)

func Register(router *mux.Router) *mux.Router {
	router.HandleFunc("/register", controllers.RegisterPost).Methods("POST")
	router.HandleFunc("/active/{user}/{activeCode}", controllers.RegisterPost).Methods("POST")

	return router
}
