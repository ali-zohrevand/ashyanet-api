package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = index(router)
	router = login(router)
	router = Register(router)
	router = Device(router)
	router = Key(router)
	router = Location(router)
	router = MqttCommand(router)
	router = Acl(router)
	router = UserDeviceLocation(router)
	router = Event(router)
	router = Active(router)
	return router
}
