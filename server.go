package main

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/routers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/ali-zohrevand/ashyanet-api/websocket"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"net/http"
)

//init
func init() {
	services.InitServices()
	go websocket.CreateWebSocketServer()
}
func main() {
	fmt.Println("Start")
	// Check if the cert files are available.
	services.GeneratTls()
	router := routers.InitRoutes()
	//corsHandler := cors.AllowAll().Handler(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		// Enable Debugging for testing, consider disabling in production
		AllowedMethods: []string{"GET", "UPDATE", "PUT", "POST", "DELETE"},
		Debug:          true,
	})
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	//err := http.ListenAndServeTLS(":5000", "cert.pem", "key.pem", n)
	err := http.ListenAndServe(":4000", n)
	fmt.Println(err)
}
