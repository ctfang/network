package main

import (
	"github.com/ctfang/network"
	"log"
)

func main() {
	server := network.NewServer("ws://127.0.0.1:8080")
	server.SetEvent(&wsserverevent{})
	server.ListenAndServe()
}

type wsserverevent struct {
}

func (*wsserverevent) OnStart(listen network.ListenTcp) {

}

func (*wsserverevent) OnConnect(connect network.Connect) {
	connect.SendString("OnConnect")
}

func (*wsserverevent) OnMessage(connect network.Connect, message []byte) {
	log.Println(string(message))
	connect.SendString("OnMessage")
}

func (*wsserverevent) OnClose(connect network.Connect) {
	log.Println("OnClose")
}

func (*wsserverevent) OnError(listen network.ListenTcp, err error) {
	log.Println("OnError")
}
