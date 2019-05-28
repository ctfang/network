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

func (c *Connect) Send(msg []byte) bool {
	sendMsg := c.Listen.Protocol().Write(msg)
	_, err := c.conn.Write(sendMsg)
	if err != nil {
		return false
	}
	return true
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
