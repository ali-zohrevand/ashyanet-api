package services

import (
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"testing"
	"time"
)

func TestCreateDevice(t *testing.T) {
	fmt.Println(time.Now())
	var TrunOnCommand models.MqttCommand
	TrunOnCommand.Name = "On"
	TrunOnCommand.Topic = "/home/root"
	TrunOnCommand.Dsc = "Turn Light On"
	TrunOnCommand.Value = "on"
	var TrunoffCommand models.MqttCommand
	TrunoffCommand.Name = "Of"
	TrunoffCommand.Topic = "/home/dffd/dfsdf"
	TrunoffCommand.Dsc = "Turn Light On"
	TrunoffCommand.Value = "off"
	var DaTa models.MqttData
	DaTa.Dsc = "status"
	DaTa.Topic = "/sds"
	DaTa.Name = "Status"
	DaTa.ValueType = "int"
	Lamp := models.DeviceInDB{}
	//FSROUPwjOKbGJjYQs5TI
	//.........................................
	Lamp.Name = "testDevice-" + GenerateRandomString(5)
	Lamp.Description = "لامپ داخل اتاقل "
	Lamp.Key = getValidKey()
	Lamp.Type = "light"
	Lamp.Location = "room"
	Lamp.Publish = append(Lamp.Publish, "sdsd/sdsad/545465/sdfsdf")
	Lamp.Publish = append(Lamp.Publish, "sddfdfsd/sdsad")
	Lamp.Publish = append(Lamp.Publish, "sdsd/sdsad/sdfsdf")
	Lamp.Subscribe = append(Lamp.Subscribe, "/dfsdfsdsd/sdsad/545465/sdfsdf")
	Lamp.Subscribe = append(Lamp.Subscribe, "sddfdfdfdfdf5456456465dfsd/sdsad")
	Lamp.Subscribe = append(Lamp.Subscribe, "/")

	Lamp.MqttCommand = append(Lamp.MqttCommand, TrunOnCommand)
	Lamp.MqttCommand = append(Lamp.MqttCommand, TrunoffCommand)
	Lamp.Pubsub = append(Lamp.Pubsub, "/d")
	Lamp.Pubsub = append(Lamp.Pubsub, "sd656456465dfsd/sdsad")
	Lamp.MqttPassword = "123456789"
	Lamp.MqttData = append(Lamp.MqttData, DaTa)
	s, _ := DB.ConnectDB()
	user, _ := DB.UserGetByUsername("ali", s)
	err, meesgat := CreateDevice(&Lamp, user)

	fmt.Println(string(meesgat), err)
	fmt.Println(time.Now())

}
