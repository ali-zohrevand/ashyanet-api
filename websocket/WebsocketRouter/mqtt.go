package WebsocketRouter

import (
	"github.com/ali-zohrevand/ashyanet-api/websocket/WebsocketHandler"
	"github.com/gorilla/mux"
)

func Mqtt(router *mux.Router) *mux.Router {
	router.HandleFunc("/data/{token}/{topic}", WebsocketHandler.MqttHandleFunc)
	return router
}
