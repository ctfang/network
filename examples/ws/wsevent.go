package main

import (
	"network"
)

type WsEvent struct {
}

func (*WsEvent) OnStart(listen network.ListenTcp) {
	println("implement me")
}

func (*WsEvent) OnConnect(connect network.Connect) {
	println("implement me")
}

func (*WsEvent) OnMessage(connect network.Connect, message interface{}) {
	println("implement me")
}

func (*WsEvent) OnClose(connect network.Connect) {
	println("implement me")
}

func (*WsEvent) OnError(listen network.ListenTcp, err error) {
	println("implement me")
}
