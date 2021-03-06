package main

import (
	"github.com/ctfang/network"
	"log"
)

func main() {
	client := network.NewClient("ws://127.0.0.1:8080")
	client.SetEvent(&clientevent{})
	client.ListenAndServe()
}

type clientevent struct {
}

func (event *clientevent) OnStart(listen network.ListenTcp) {

}

func (*clientevent) OnConnect(connect network.Connect) {
}

func (*clientevent) OnMessage(connect network.Connect, message []byte) {
	log.Println(string(message))
	connect.Send("[123]")
}

func (*clientevent) OnClose(connect network.Connect) {

}

func (*clientevent) OnError(listen network.ListenTcp, err error) {

}
