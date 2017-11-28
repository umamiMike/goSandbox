package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8090", nil)
	fmt.Println("Serving on Port 8090")
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
