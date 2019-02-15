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
	fmt.Println("connections created")
	// Make sure we close the connection when the function returns
	defer ws.Close()

	vars := mux.Vars(r)
	var token, topic string
	token = vars["token"]
	topic = vars["topic"]
	if token == "" || topic == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
	}
	status, _ := services.IsJwtValid(token)
	if status != http.StatusOK {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(""))
		return
	}
	mqttObject, errSubscribe := services.MqttSubcribeTopicWebsocket(topic, ws)
	if errSubscribe != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	if mqttObject.Client != nil && errSubscribe == nil {
		defer mqttObject.Client.Unsubscribe(topic)

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
