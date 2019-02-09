package main

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/routers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"net/http"
)

//init
func init() {
	services.InitServices()
}
func main() {
	// Check if the cert files are available.
	services.GeneratTls()
	router := routers.InitRoutes()
	corsHandler := cors.AllowAll().Handler(router)

	n := negroni.Classic()
	n.UseHandler(corsHandler)
	err := http.ListenAndServeTLS(":5000", "cert.pem", "key.pem", n)
	//err := http.ListenAndServe(":5000",  n)
	fmt.Println(err)
}
