package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Person struct {
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Email string `json:"email"`
	About string `json:"about"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Client Connected")
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
