package tcp

import (
	"github.com/ctfang/network"
	"net"
)

type ListenTcp struct {
	id         uint32
	url        *network.Url
	event      network.Event
	protocol   network.Protocol
	conn       net.Conn
	newConnect func(listen network.ListenTcp, conn net.Conn) network.Connect
}

func (c *ListenTcp) SetUrl(address *network.Url) {
	c.url = address
}

func (c *ListenTcp) Url() *network.Url {
	return c.url
}

func (c *ListenTcp) SetEvent(event network.Event) {
	c.event = event
}

func (c *ListenTcp) Event() network.Event {
	return c.event
}

func (c *ListenTcp) SetProtocol(protocol network.Protocol) {
	c.protocol = protocol
}

func (c *ListenTcp) Protocol() network.Protocol {
	return c.protocol
}

func (c *ListenTcp) Close() {
	_ = c.conn.Close()
}

func (c *ListenTcp) SetNewConnect(new func(listen network.ListenTcp, conn net.Conn) network.Connect) {
	c.newConnect = new
}
