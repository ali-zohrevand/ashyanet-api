package main

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/Lab/Emqtt/DB"
	"github.com/ali-zohrevand/ashyanet-api/Lab/Emqtt/models"
)

func main() {
	mqttuser := models.MqttUser{}
	mqttuser.Password = "123456"
	mqttuser.Username = "root"
	mqttuser.Is_superuser = true
	session, _ := DB.ConnectDB()
	err := DB.CreateMqttUser(mqttuser, session)
	fmt.Println(err)
}
