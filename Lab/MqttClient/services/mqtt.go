package services

import "github.com/eclipse/paho.mqtt.golang"

type mqttStruct struct {
	Ip       string
	UserName string
	Password string
	option   mqtt.ClientOptions
}

func NewMqtt(ip string, username string, passwoard string) *mqttStruct {
	return &mqttStruct{ip, username, passwoard}
}

func (ds *mqttStruct) Publish() {

}
