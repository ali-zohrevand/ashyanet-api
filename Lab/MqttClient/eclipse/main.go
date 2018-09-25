package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

var a mqtt.MessageHandler

//https://www.hivemq.com/blog/mqtt-client-library-encyclopedia-golang
func main() {
	done := make(chan bool)
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("sample")
	opts.SetUsername("root")
	opts.SetPassword("123456")
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	//Send Message
	c.Publish("test/topic", 1, true, "Example Payload")
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())

	}
	token := c.Subscribe("#", 0, f)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	<-done
}
