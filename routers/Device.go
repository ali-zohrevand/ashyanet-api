package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Device(router *mux.Router) *mux.Router {
	//.......................................................................................................
	HandleFunc := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.CreateDevice),
	)
	router.Handle("/user/devices", HandleFunc).Methods("POST")
	//.......................................................................................................

	HandleFuncList := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Index),
	)
	router.Handle("/user/devices", HandleFuncList).Methods("GET")
	//.......................................................................................................
	HandleFuncAddUserToDevice := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.AddUserToDevice),
	)
	router.Handle("/user/aud", HandleFuncAddUserToDevice).Methods("POST")
	//.......................................................................................................
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
