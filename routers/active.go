package routers

import (
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/controllers"
)

func Active(router *mux.Router) *mux.Router {
	router.HandleFunc("/active/{user}/{activeCode}", controllers.Active).Methods("GET")

	return router
}
