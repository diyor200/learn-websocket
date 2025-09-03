package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	go startClient()
	http.HandleFunc("/echo", echo)

	http.ListenAndServe(":8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("received: %s\n", string(message))

		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
