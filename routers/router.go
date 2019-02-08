package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = index(router)
	router = User(router)
	router = Device(router)
	router = Key(router)
	router = Location(router)
	router = MqttCommand(router)
	router = Acl(router)
	router = UserDeviceLocation(router)
	router = Event(router)
	router = Active(router)
	router = Types(router)
	router = Jwt(router)
	router = Info(router)
	return router
}
