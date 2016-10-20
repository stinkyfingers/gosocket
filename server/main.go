package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/buff", websocket.Handler(handlerBuffer))
	http.Handle("/", websocket.Handler(handler))
	log.Fatal(http.ListenAndServe(":7000", nil))
}

// encoded, any len
func handler(ws *websocket.Conn) {
	type Data struct {
		Message string
	}
	for {
		var d Data
		err := websocket.JSON.Receive(ws, &d)
		if err != nil {
			log.Print(err)
			break
		}
		log.Print(d)
	}
}

// limited to 1024 bytes
func handlerBuffer(ws *websocket.Conn) {
	// make a buffer
	var buf = make([]byte, 1024)
	n, err := ws.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("We read", n, "bytes, and they were", string(buf))

}
