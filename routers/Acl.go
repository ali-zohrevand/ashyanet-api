package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Acl(router *mux.Router) *mux.Router {
	//.......................................................................................................
	HandleFunc := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.CreateAcl),
	)
	router.Handle("/user/acl", HandleFunc).Methods("POST")
	//.......................................................................................................

	return router
}
