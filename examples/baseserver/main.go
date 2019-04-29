package main

import (
	"network"
	"network/tcp"
)

func main()  {
	server := tcp.Server{}
	server.SetAddress(network.NewAddress("127.0.0.1:8080"))
	server.SetConnectionEvent(NewBaseEvent())
	server.ListenAndServe()
}