package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/controllers"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
)

func Device(router *mux.Router) *mux.Router {
	//.......................................................................................................
	HandleFunc := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.CreateDevice),
	)
	router.Handle("/user/devvices", HandleFunc).Methods("POST")
	//.......................................................................................................

	HandleFuncList := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Index),
	)
	router.Handle("/user/devices", HandleFuncList).Methods("GET")
	//.......................................................................................................

	HandleFuncAddKeyToDevice := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.AddKeyToDevice),
	)
	router.Handle("/user/akd", HandleFuncAddKeyToDevice).Methods("POST")
	//.......................................................................................................
	//.......................................................................................................

	return router
}
