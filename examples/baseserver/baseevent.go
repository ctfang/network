package main

import (
	"github.com/ctfang/network"
	"log"
)

type baseevent struct {
}

func (*baseevent) OnStart(listen network.ListenTcp) {
	log.Println("OK")
}

func (*baseevent) OnConnect(connect network.Connect) {
	log.Println("OK")
}

func (*baseevent) OnMessage(connect network.Connect, message interface{}) {
	log.Println("OK")
}

func (*baseevent) OnClose(connect network.Connect) {
	log.Println("OK")
}

func (*baseevent) OnError(listen network.ListenTcp, err error) {
	log.Println("OK")
}

func NewBaseEvent() *baseevent {
	return &baseevent{}
}
