package controllers

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services"
	"github.com/gorilla/mux"
	"net/http"
)

func MqttTopicsGetAll(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.MqttGetAllTopicsByUsername(userInDB.UserName, "all")
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
func MqttTopicsGetAllByType(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	typeOfTopics := vars["type"]
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.MqttGetAllTopicsByUsername(userInDB.UserName, typeOfTopics)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
func MqttTopicGetData(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	mqttInfo := new(models.MqttDataRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&mqttInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.MqttGetALlMessageByTopicTypeUsername(userInDB.UserName, mqttInfo.Topic)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
func MqttTopicGetDataInfo(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userInDB, err := GetUserFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))

	} else {
		w.Header().Set("Content-Type", "application/json")
		responseStatus, token := services.MqttGetALlInfoTypeUsername(userInDB.UserName)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}
