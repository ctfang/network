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

func (c *ConnectTcp) Send(msg interface{}) bool {
	panic("implement me")
}

func (c *ConnectTcp) Url() *network.Url {
	return c.url
}
