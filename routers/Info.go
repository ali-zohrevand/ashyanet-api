package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Info(router *mux.Router) *mux.Router {
	HandleFunc := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Info),
	)
	router.Handle("/user/info", HandleFunc).Methods("GET")
	return router
}
