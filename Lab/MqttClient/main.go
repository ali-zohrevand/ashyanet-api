package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"gitlab.com/hooshyar/ChiChiNi-API/Lab/MqttClient/services"
	"time"
)

func main() {
	done := make(chan bool)

	m, er := services.NewMqtt("tcp://127.0.0.1:1883", "bob", "123456", "gateway")
	fmt.Println(er)
	err := m.Publish("/World", false, "sadddlam", 2)
	fmt.Println(err)

	go func() {
		done := make(chan bool)
		var eventFunc mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
			fmt.Println(message.Topic(), "  ", string(message.Payload()), " d ", message.Duplicate())
		}
		m.Subscribe("#", 1, eventFunc)
		<-done

	}()
	for {
		fmt.Println("Time is: ", time.Now())
		time.Sleep(time.Second * 10)
	}
	<-done

}
