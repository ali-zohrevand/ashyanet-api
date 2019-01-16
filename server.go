package main

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/routers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"net/http"
)

func init() {
	services.InitServices()
}
func main() {
	// Check if the cert files are available.
	services.GeneratTls()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	err := http.ListenAndServeTLS(":5000", "cert.pem", "key.pem", n)
	fmt.Println(err)
}
