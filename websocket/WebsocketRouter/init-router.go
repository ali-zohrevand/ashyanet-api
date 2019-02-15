package WebsocketRouter

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = Mqtt(router)
	return router
}
