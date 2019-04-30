package main

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/tcp"
)

func main() {
	server := tcp.Server{}
	server.SetAddress(network.NewAddress("127.0.0.1:8080"))
	server.SetConnectionEvent(&baseevent{})
	server.ListenAndServe()
}

type baseevent struct {
}

func (*baseevent) OnStart(listen network.ListenTcp) {
	println("implement me")
}

func (*baseevent) OnConnect(connect network.Connect) {
	println("implement me")
}

func (*baseevent) OnMessage(connect network.Connect, message interface{}) {
	println("implement me")
}

func (*baseevent) OnClose(connect network.Connect) {
	println("implement me")
}

func (*baseevent) OnError(listen network.ListenTcp, err error) {
	println("implement me")
}
