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
		negroni.HandlerFunc(controllers.DeviceCreate),
	)
	router.Handle("/user/devices", HandleFunc).Methods("POST")
	//.......................................................................................................

	HandleFuncList := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.DeviceGetAll),
	)
	router.Handle("/user/devices", HandleFuncList).Methods("GET")
	//.......................................................................................................
	HandleFuncGetId := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.DeviceGetId),
	)
	router.Handle("/user/devices/{id}", HandleFuncGetId).Methods("GET")
	//.......................................................................................................
	HandleFuncDeleteId := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.DeviceDeleteId),
	)
	router.Handle("/user/devices/{id}", HandleFuncDeleteId).Methods("DELETE")
	//.......................................................................................................
	HandleFuncUpdateId := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.DeviceUpdateId),
	)
	router.Handle("/user/devices/{id}", HandleFuncUpdateId).Methods("PUT")
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
