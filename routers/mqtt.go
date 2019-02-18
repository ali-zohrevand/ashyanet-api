package routers

import (
	"github.com/ali-zohrevand/ashyanet-api/controllers"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Mqtt(router *mux.Router) *mux.Router {

	HandleFuncMqttTopicsGetAll := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.MqttTopicsGetAll),
	)
	router.Handle("/user/mqtt", HandleFuncMqttTopicsGetAll).Methods("GET")
	//.....................................................................................................
	HandleFuncMqttTopicsGetAllByType := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.MqttTopicsGetAllByType),
	)
	router.Handle("/user/mqtt/{type}", HandleFuncMqttTopicsGetAllByType).Methods("GET")
	//.....................................................................................................
	HandleFunfMqttGetData := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.MqttTopicGetData),
	)
	router.Handle("/user/data", HandleFunfMqttGetData).Methods("POST")
	//.....................................................................................................
	HandleFunfMqttGetDataInfo := negroni.New(
		negroni.HandlerFunc(services.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.MqttTopicGetDataInfo),
	)
	router.Handle("/user/data", HandleFunfMqttGetDataInfo).Methods("GET")

	return router
}
