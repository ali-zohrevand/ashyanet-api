package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/controllers"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
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
