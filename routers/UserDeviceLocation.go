package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func UserDeviceLocation(router *mux.Router) *mux.Router {
	HandleFuncAddUserToDevice := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.AddUserToDevice),
	)
	router.Handle("/user/aud", HandleFuncAddUserToDevice).Methods("POST")
	//.......................................................................................................
	HandleFuncAddUserToLocation := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.AddUserToLocation),
	)
	router.Handle("/user/aul", HandleFuncAddUserToLocation).Methods("POST")
	//.......................................................................................................
	return router

}
