package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Location(router *mux.Router) *mux.Router {
	HandleFuncCreateLocation := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.LocationCreate),
	)
	router.Handle("/user/locations", HandleFuncCreateLocation).Methods("POST")
	//.......................................................................................................
	HandleFuncGetAllLocation := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.LocationGetAll),
	)
	router.Handle("/user/locations", HandleFuncGetAllLocation).Methods("GET")
	//.......................................................................................................
	HandleFuncDeleteLocation := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.LocationDeleteById),
	)
	router.Handle("/user/locations/{id}", HandleFuncDeleteLocation).Methods("DELETE")

	return router

}
