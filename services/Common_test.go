package services

import (
	"encoding/json"
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"testing"
)

func TestAddRootTopic(t *testing.T) {
	p := AddRootTopic("sadasdasd", "/")
	fmt.Println(p)
}
func TestCheckMqttTopic(t *testing.T) {
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
	Lamp.Name = "lamp"
	Lamp.Description = "لامپ داخل اتاقل "
	Lamp.Key = getValidKey()
	Lamp.Type = "light"
	Lamp.Location = "room"
	Lamp.Publish = append(Lamp.Publish, "sdsd/sdsad/545465/sdfsdf")
	Lamp.Publish = append(Lamp.Publish, "sddfdfsd/sdsad")
	Lamp.Publish = append(Lamp.Publish, "sdsd/sdsad/sdfsdf")
	Lamp.Subscribe = append(Lamp.Subscribe, "/dfsdfsdsd/sdsad/545465/sdfsdf")
	Lamp.Subscribe = append(Lamp.Subscribe, "sddfdfdfdfdf5456456465dfsd/sdsad")
	Lamp.Subscribe = append(Lamp.Subscribe, "")
	Lamp.Subscribe = append(Lamp.Subscribe, "/")
	Lamp.MqttCommand = append(Lamp.MqttCommand, TrunOnCommand)
	Lamp.MqttCommand = append(Lamp.MqttCommand, TrunoffCommand)
	Lamp.Pubsub = append(Lamp.Pubsub, "/d")
	Lamp.Pubsub = append(Lamp.Pubsub, "sd656456465dfsd/sdsad")
	Lamp.MqttData = append(Lamp.MqttData, DaTa)
	s, _ := DB.ConnectDB()
	user, e := DB.UserGetByUsername("ali", s)
	fmt.Println(e)
	fmt.Println("=========\n", user)
	newUser, _ := CheckMqttTopic(&Lamp, user)
	j, _ := json.Marshal(newUser)
	fmt.Println(string(j))
}
