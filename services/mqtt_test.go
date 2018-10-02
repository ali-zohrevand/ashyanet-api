package services

import (
	"encoding/json"
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"testing"
	"time"
)

func TestMqttAddMessage(t *testing.T) {

	var m models.MqttMessage
	m.Message = "test"
	m.Qos = 2
	m.Retained = false
	m.Topic = "/test"
	m.Time = time.Now().String()
	err := MqttAddMessage(m)
	if err != nil {
		t.Fail()
		t.Error(err)
	}

}
func TestMqttGetAllMessageByTopicName(t *testing.T) {
	list, err := MqttGetAllMessageByTopicName("/test")
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	mesage := list[0]
	json, _ := json.Marshal(mesage)
	fmt.Println(string(json))

}