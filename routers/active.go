package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/gorilla/mux"
)

func Active(router *mux.Router) *mux.Router {
	router.HandleFunc("/active/{user}/{activeCode}", controllers.Active).Methods("GET")

	return router
}
