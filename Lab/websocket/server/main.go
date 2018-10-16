package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

type message struct {
	// the json tag means this will serialize as a lowercased field
	Message string `json:"message"`
}

func socket(ws *websocket.Conn) {
	for {
		// allocate our container struct
		var m message
		m2 := message{"Thanks for the message!"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}
		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}

		log.Println("Received message:", m.Message)

		// send a response
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}
	}
}
func main() {
	// later, in your router initialization:
	http.Handle("/socket", websocket.Handler(socket))
	fmt.Print(http.ListenAndServe(":5000", nil))
}
