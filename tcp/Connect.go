package tcp

import (
	"github.com/ctfang/network"
	"net"
)

var id uint32

type Connect struct {
	ConnectTcp
}

func (c *Connect) GetIp() uint32 {
	panic("implement me")
}

func (c *Connect) GetPort() uint16 {
	panic("implement me")
}

func NewConnect(listen network.ListenTcp, conn net.Conn) network.Connect {
	id = id + 1
	url := network.NewUrl(listen.Url().Scheme + "://" + listen.Url().Host)
	return &Connect{
		ConnectTcp: ConnectTcp{
			id:     id,
			uid:    "",
			url:    url,
			conn:   conn,
			Listen: listen,
		},
	}
}