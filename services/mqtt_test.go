package services

import (
	"encoding/json"
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
	"time"
)

func TestMqttAddMessage(t *testing.T) {

	var m models.MqttMessage
	m.Message = "test"
	m.Qos = 2
	m.Retained = false
	m.Topic = "/test"
	m.Time = time.Now().Unix()
	err := MqttAddMessageToDb(m)
	if err != nil {
		t.Fail()
		t.Error(err)
	}

}
func TestMqttGetAllMessageByTopicName(t *testing.T) {
	list, err := MqttGetAllMessageByTopicName("/test")
	if err != nil || list == nil {
		t.Fail()
		t.Error(err)
	} else {
		mesage := list[0]
		json, _ := json.Marshal(mesage)
		fmt.Println(string(json))
	}

}
func TestMqttGetAllTopicsByUsername(t *testing.T) {
	list, err := MqttGetAllTopicsByUsername("mahmood", "sub")
	fmt.Println(list, err)
}
