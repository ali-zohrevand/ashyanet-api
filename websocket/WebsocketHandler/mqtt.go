package WebsocketHandler

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

var ws *websocket.Conn

func MqttHandleFunc(w http.ResponseWriter, r *http.Request) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}

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
	userIndb, err := DB.JwtGetUser(token, session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	fmt.Println(userIndb)
	HasUser, err := DB.UserHasTopic(topic, userIndb.UserName, "sub", session)
	if !HasUser || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(""))
		return
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	fmt.Println("connections created")
	// Make sure we close the connection when the function returns
	defer ws.Close()
	go ChecKwebSockerStatu(ws)
	mqttObject, errSubscribe := services.MqttSubcribeTopicWebsocket(topic, ws)
	if errSubscribe != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	defer mqttObject.Client.Unsubscribe(topic)
	defer mqttObject.Client.Disconnect(50)

}
func ChecKwebSockerStatu(webscoketObj *websocket.Conn) {

}
