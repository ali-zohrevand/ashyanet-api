// +build mage
package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"gitlab.com/hooshyar/ChiChiNi-API/routers"
	"gitlab.com/hooshyar/ChiChiNi-API/services"
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
