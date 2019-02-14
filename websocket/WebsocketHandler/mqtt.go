package WebsocketHandler

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var ws *websocket.Conn

func MqttHandleFunc(w http.ResponseWriter, r *http.Request) {
	userIndb, err := controllers.GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	fmt.Println(userIndb)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// Upgrade initial GET request to a websocket
	ws, err = upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	vars := mux.Vars(r)
	token := vars["token"]
	status, _ := services.IsJwtValid(token)
	if status != http.StatusOK {

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(""))
		return
	}
}
func MqttTopicSubscribe() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
