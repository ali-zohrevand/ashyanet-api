package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gitlab.com/hooshyar/ChiChiNi-API/controllers"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
)

func index(router *mux.Router) *mux.Router {
	HandleFunc := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Index),
	)
	router.Handle("/", HandleFunc).Methods("GET")
	return router
}

/*
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
*/