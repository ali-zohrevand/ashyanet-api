package websocket

import (
	"github.com/ali-zohrevand/ashyanet-api/websocket/WebsocketRouter"
	"log"
	"net/http"
)

func CreateWebSocketServer() {
	router := WebsocketRouter.InitRoutes()
	//......................................
	//http.HandleFunc("/ws", handleConnections)
	log.Println("http server started on :1234")
	err := http.ListenAndServe(":1234", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
