package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"os"
)

const address string = "localhost:5000"

func main() {

	initWebsocketClient()
}

type messageClass struct {
	message string
}

func initWebsocketClient() {
	fmt.Println("Starting Client")
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/socket", address), "", fmt.Sprintf("http://%s/", address))
	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		os.Exit(1)
	}
	incomingMessages := make(chan string)
	go readClientMessages(ws, incomingMessages)
	var i = 0
	for {
		i++
		select {

		case message := <-incomingMessages:
			fmt.Println(`Message Received:`, message)
			var m messageClass

			m.message = "salam" + string(i)

			j, _ := json.Marshal(m)
			websocket.JSON.Send(ws, j)
		default:
			continue

		}
	}
}

func readClientMessages(ws *websocket.Conn, incomingMessages chan string) {
	for {
		var message string
		// err := websocket.JSON.Receive(ws, &message)
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Printf("Error::: %s\n", err.Error())
			return
		}
		incomingMessages <- message
	}
}
