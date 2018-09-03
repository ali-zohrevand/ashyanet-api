package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/controllers"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
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
