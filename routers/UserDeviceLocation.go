package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/controllers"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
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
