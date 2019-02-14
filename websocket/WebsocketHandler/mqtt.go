package WebsocketHandler

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/gorilla/mux"
	"net/http"
)

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
	ws, err := upgrader.Upgrade(w, r, nil)

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
