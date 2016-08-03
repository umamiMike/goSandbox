package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

var (
	dirPath string
	test    byte
	lenFile int
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ChatServer(ws *websocket.Conn) {
	defer ws.Close()
	var test []byte
	var payload []byte
	for {
		err := websocket.Message.Receive(ws, &payload)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Payload: ", len(payload))

		test = append(test, payload...)
		fmt.Println("Test: ", len(test))

		fo, err := os.Create(fmt.Sprintf("./%d.png", time.Now().UnixNano()))
		check(err)
		_, err = fo.Write(test)
		check(err)
		fo.Close()
	}
	log.Print("DONE")

}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: chatExample <dir>")
	}

	dirPath = os.Args[1]

	fmt.Println("Starting... ")

	http.Handle("/", http.FileServer(http.Dir(dirPath)))
	http.Handle("/ws", websocket.Handler(ChatServer))

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}
