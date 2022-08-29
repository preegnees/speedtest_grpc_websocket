package main

import (
	"log"
	"net/http"

	websocket "github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{}

func saveFile(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Connect err: ", err)
	}
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			break
		}
		log.Printf("messageType: %d, message: %s", messageType, message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("write err:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/", saveFile)
	log.Fatal(http.ListenAndServe(":55001", nil))
}
