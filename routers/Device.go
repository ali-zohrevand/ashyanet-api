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
	router.Handle("/devices", HandleFunc).Methods("POST")
	//.......................................................................................................

	HandleFuncList := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Index),
	)
	router.Handle("/devices", HandleFuncList).Methods("GET")
	//.......................................................................................................
	HandleFuncAddUserToDevice := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.AddUserToDevice),
	)
	router.Handle("/aud", HandleFuncAddUserToDevice).Methods("POST")
	//.......................................................................................................
	//.......................................................................................................
	HandleFuncAddKeyToDevice := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.AddKeyToDevice),
	)
	router.Handle("/akd", HandleFuncAddKeyToDevice).Methods("POST")
	//.......................................................................................................
	//.......................................................................................................

	return router
}
