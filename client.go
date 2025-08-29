package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func startClient() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Println("Connecting to", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Handle Ping automatically by replying with Pong
	conn.SetPingHandler(func(appData string) error {
		fmt.Println("Got Ping:", appData)
		return conn.WriteMessage(websocket.PongMessage, []byte(appData))
	})

	// Keep reading to keep connection alive
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		fmt.Println("Received message:", string(msg))
	}
}
