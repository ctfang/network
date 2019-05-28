package main

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/tool"
	"log"
)

func main() {
	client := tool.NewServer("ws://127.0.0.1:8080")
	client.SetEvent(&wsserverevent{})
	client.ListenAndServe()
}

type wsserverevent struct {
}

func (*wsserverevent) OnStart(listen network.ListenTcp) {

}

func (*wsserverevent) OnConnect(connect network.Connect) {

}

func (*wsserverevent) OnMessage(connect network.Connect, message []byte) {
	log.Println(string(message))
}

func (*wsserverevent) OnClose(connect network.Connect) {
	log.Println("OnClose")
}

func (*wsserverevent) OnError(listen network.ListenTcp, err error) {
	log.Println("OnError")
}
