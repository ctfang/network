package tool

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/protocol"
	"github.com/ctfang/network/tcp"
)

func NewClient(address string) network.ListenTcp {
	client := tcp.Client{}
	client.SetProtocol(&protocol.WebsocketProtocol{})
	client.SetUrl(network.NewUrl(address))
	return &client
}
