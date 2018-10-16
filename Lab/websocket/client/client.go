package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"os"
)

const address string = "localhost:5000"

func main() {

	initWebsocketClient()
}

type message struct {
	// the json tag means this will serialize as a lowercased field
	Message string `json:"message"`
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

		case messageRecived := <-incomingMessages:
			fmt.Println(`Message Received:`, messageRecived)
			var m message

			m.Message = "salam " + string(i)

			websocket.JSON.Send(ws, m)
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
