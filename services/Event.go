package services

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
)

func MqttSubcribeRootTopic() (err error) {
	EmqttDeleteMqttDefaultAdmin()
	TempUserAdminUserName, TempAdminPassword, errCreateTempAdmin := EmqttCreateTempAdminMqttUserWithDwefaultAdmin()
	if errCreateTempAdmin != nil {
		panic(err)
		return errCreateTempAdmin
	}
	defer EmqttDeleteUser(TempUserAdminUserName)
	done := make(chan bool)
	mqttObj, errCreateMqttUser := NewMqtt(Words.MqttBrokerIp, TempUserAdminUserName, TempAdminPassword, "TempAdmin")
	if errCreateMqttUser != nil {
		panic(err)
		return errCreateMqttUser
	}
	var eventFunc mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
		fmt.Println(message.Topic(), "  ", string(message.Payload()))
	}
	errSubscribe := mqttObj.Subscribe("#", 2, eventFunc)
	if errSubscribe != nil {
		panic(err)
		return errSubscribe
	}
	<-done
	return
}
