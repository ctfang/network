package tcp

import (
	"github.com/ctfang/network"
	"net"
)

type ConnectTcp struct {
	id     uint32
	uid    string
	url    *network.Url
	conn   net.Conn
	Listen network.ListenTcp
	header network.Header
}

func (c *ConnectTcp) GetIp() uint32 {
	panic("implement me")
}

func (c *ConnectTcp) GetPort() uint16 {
	return c.url.Port
}

func (c *ConnectTcp) GetCon() net.Conn {
	return c.conn
}

func (c *ConnectTcp) Close() {
	_ = c.conn.Close()
}

func (c *ConnectTcp) Id() uint32 {
	return c.id
}

func (c *ConnectTcp) SetUid(uid string) {
	c.uid = uid
}

func (c *ConnectTcp) Uid() string {
	return c.uid
}

func (c *ConnectTcp) Url() *network.Url {
	return c.url
}

func (c *ConnectTcp) SendByte(msg []byte) bool {
	err := c.Listen.Protocol().Write(c.conn, msg)
	if err != nil {
		return false
	}
	return true
}

func (c *ConnectTcp) SendString(msg string) bool {
	err := c.Listen.Protocol().Write(c.conn, []byte(msg))
	if err != nil {
		return false
	}
	return true
}

func (c *ConnectTcp) SetHeader(header network.Header) {
	c.header = header
}
func (c *ConnectTcp) Header() network.Header {
	return c.header
}
