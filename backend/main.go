package main

import (
	"fmt"
	// "net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Hello!")
	fmt.Println(mux.NewRouter())
	fmt.Println(websocket.PingMessage)
}
