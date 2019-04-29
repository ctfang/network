package main

import (
	"network"
	"network/ws"
)

func main() {
	server := ws.Server{}
	server.SetAddress(network.NewAddress("127.0.0.1:8081"))
	server.SetConnectionEvent(&WsEvent{})
	server.ListenAndServe()
}
