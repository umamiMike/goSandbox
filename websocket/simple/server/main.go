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
	http.Handle("/", http.FileServer(http.Dir("/home/mike/dev/wilding/goSandbox/websocket/simple/web client")))
	err := http.ListenAndServe(":5060", nil)
	fmt.Println("Serving on Port 5060")
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
