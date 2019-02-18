package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Types(router *mux.Router) *mux.Router {
	HandleFuncTypesCreate := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.TypesCreate),
	)
	HandleFuncTypesGetAll := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.TypesGetAll),
	)
	HandleFuncTypesDelete := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.TypesDelete),
	)
	router.Handle("/user/types", HandleFuncTypesCreate).Methods("POST")
	router.Handle("/user/types", HandleFuncTypesGetAll).Methods("GET")
	router.Handle("/user/types/{id}", HandleFuncTypesDelete).Methods("DELETE")

	//.......................................................................................................
	return router

}
