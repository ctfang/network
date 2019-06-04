package network

import (
	"net"
)

type Listen struct {
	id         uint32
	url        *Url
	event      Event
	protocol   Protocol
	conn       net.Conn
	newConnect func(listen ListenTcp, conn net.Conn) Connect
}

func (c *Listen) SetUrl(address *Url) {
	c.url = address
}

func (c *Listen) Url() *Url {
	return c.url
}

func (c *Listen) SetEvent(event Event) {
	c.event = event
}

func (c *Listen) Event() Event {
	return c.event
}

func (c *Listen) SetProtocol(protocol Protocol) {
	c.protocol = protocol
}

func (c *Listen) Protocol() Protocol {
	return c.protocol
}

func (c *Listen) Close() {
	_ = c.conn.Close()
}

func (c *Listen) SetNewConnect(new func(listen ListenTcp, conn net.Conn) Connect) {
	c.newConnect = new
}
