package tool

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/protocol"
	"github.com/ctfang/network/tcp"
)

func NewClient(address string) network.ListenTcp {
	url := network.NewUrl(address)
	client := tcp.Client{}
	switch url.Scheme {
	case "ws":
		client.SetProtocol(&protocol.WsProtocol{})
	case "text":
		client.SetProtocol(&protocol.TextProtocol{})
	default:
		panic("ws or text")
	}
	client.SetUrl(url)
	return &client
}

func NewServer(address string) network.ListenTcp {
	url := network.NewUrl(address)

	server := tcp.Server{}
	switch url.Scheme {
	case "ws":
		server.SetProtocol(&protocol.WebsocketProtocol{})
	case "text":
		server.SetProtocol(&protocol.TextProtocol{})
	default:
		panic("ws or text")
	}

	server.SetUrl(url)
	return &server
}
