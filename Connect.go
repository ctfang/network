package network

import (
	"net"
)

var id uint32

type Connection struct {
	ConnectTcp
}

func (c *Connection) Send(msg interface{}) bool {
	var err error
	switch msg.(type) {
	case []byte:
		return c.SendByte(msg.([]byte))
	case string:
		return c.SendString(msg.(string))
	}

	if err != nil {
		return false
	}
	return true
}

func NewConnect(listen ListenTcp, conn net.Conn) Connect {
	id = id + 1
	url := NewUrl(listen.Url().Scheme + "://" + listen.Url().Host)
	return &Connection{
		ConnectTcp: ConnectTcp{
			id:     id,
			uid:    "",
			url:    url,
			conn:   conn,
			Listen: listen,
		},
	}
}
