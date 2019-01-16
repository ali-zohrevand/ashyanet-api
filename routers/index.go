package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func index(router *mux.Router) *mux.Router {
	HandleFunc := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Index),
	)
	UpFunc := negroni.New(
		negroni.HandlerFunc(controllers.Status),
	)
	router.Handle("/", HandleFunc).Methods("GET")
	router.Handle("/status", UpFunc).Methods("GET")

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
