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
		negroni.HandlerFunc(controllers.CreateLocation),
	)
	router.Handle("/user/locations", HandleFuncCreateLocation).Methods("POST")
	//.......................................................................................................
	return router

}
