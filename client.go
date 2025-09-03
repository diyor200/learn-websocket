package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func startClient() {
	u := url.URL{Scheme: "ws", Path: "/echo", Host: "localhost:8080"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%s\n", string(message))
		}
	}()

	for {
		err = conn.WriteMessage(websocket.TextMessage, []byte("hello"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
