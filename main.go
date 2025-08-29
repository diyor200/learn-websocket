package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	conn.SetPongHandler(func(appData string) error {
		log.Printf("pong: %s", appData)
		return nil
	})

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteMessage(websocket.PingMessage, []byte("are you alive?")); err != nil {
					log.Println("❌ Write ping error:", err)
					return // only exit if connection is broken
				}
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("❌ Read error:", err)
			return
		}
		fmt.Println("📩 Server received message:", string(msg))
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	fmt.Println("Server started at :8080")
	go startClient()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
