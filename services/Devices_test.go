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
	var TrunOnCommand models.Command
	TrunOnCommand.Name = "On"
	TrunOnCommand.Topic = "/home/root"
	TrunOnCommand.Dsc = "Turn Light On"
	TrunOnCommand.Value = "on"
	var TrunoffCommand models.Command
	TrunoffCommand.Name = "Of"
	TrunoffCommand.Topic = "/home/dffd/dfsdf"
	TrunoffCommand.Dsc = "Turn Light On"
	TrunoffCommand.Value = "off"
	var DaTa models.Data
	DaTa.Dsc = "status"
	DaTa.Topic = "/sds"
	DaTa.Name = "Status"
	DaTa.ValueType = "int"
	Lamp := models.Device{}
	//FSROUPwjOKbGJjYQs5TI
	//.........................................
	/*	Lamp.Name = "testDevice-" + GenerateRandomString(5)
		Lamp.Description = "لامپ داخل اتاقل "
		Lamp.Key = getValidKey()
		Lamp.Type = "light"
		Lamp.Location = "room"
		Lamp.Publish = append(Lamp.Publish, "sdsd/sdsad/545465/sdfsdf")
		Lamp.Publish = append(Lamp.Publish, "sddfdfsd/sdsad")
		Lamp.Publish = append(Lamp.Publish, "sdsd/sdsad/sdfsdf")
		Lamp.Subscribe = append(Lamp.Subscribe, "/dfsdfsdsd/sdsad/545465/sdfsdf")
		Lamp.Subscribe = append(Lamp.Subscribe, "sddfdfdfdfdf5456456465dfsd/sdsad")
		Lamp.Subscribe = append(Lamp.Subscribe, "/")*/

	//.......................................................
	Lamp.Name = "SerialSensor"
	Lamp.Description = "سنسور سریال"
	Lamp.Key = getValidKey()
	Lamp.Type = "sensor"
	Lamp.Location = "room"
	Lamp.Publish = append(Lamp.Publish, "/sensor")
	Lamp.Subscribe = append(Lamp.Subscribe, "/sensor")
	Lamp.MqttPassword = "123456789"
	Lamp.MqttData = append(Lamp.MqttData, DaTa)
	s, _ := DB.ConnectDB()
	user, _ := DB.UserGetByUsername("ali", s)
	err, meesgat := CreateDevice(&Lamp, user)

	fmt.Println(string(meesgat), err)
	fmt.Println(time.Now())

}
