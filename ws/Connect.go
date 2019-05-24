package ws

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/tcp"
	"net"
)

func NewConnect(listen network.ListenTcp, conn net.Conn) network.Connect {
	connect := tcp.NewConnect(listen, conn)

	return connect
}
