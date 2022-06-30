package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		_, m, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(string(m))
	})
	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)
	log.Println("Server running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
