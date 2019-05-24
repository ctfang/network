package tool

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/tcp"
)

func NewClient(address string) network.ListenTcp {
	client := tcp.Client{}
	client.SetUrl(network.NewUrl(address))
	return &client
}
