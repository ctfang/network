package tcp

import (
	"github.com/ctfang/network"
	"net"
	"regexp"
	"strconv"
)

type Connection struct {
	cid    uint32
	uid    string
	ip     uint32
	port   uint16
	con    net.Conn
	pro    network.Protocol
	extend interface{}
}

func (c *Connection) SetExtend(extend interface{}) {
	c.extend = extend
}

func (c *Connection) GetExtend() interface{} {
	return c.extend
}

func (c *Connection) GetIp() uint32 {
	if c.ip != 0 {
		return c.ip
	}
	ipStr := c.con.RemoteAddr().String()
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return 0
	}
	ips := reg.FindStringSubmatch(ipStr)
	if ips == nil {
		return 0
	}

	c.ip = network.Ip2long(ips[0])
	return c.ip
}

func (c *Connection) GetPort() uint16 {
	if c.port != 0 {
		return c.port
	}

	ipStr := c.con.RemoteAddr().String()
	r := `\:(\d{1,5})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return 0
	}
	ips := reg.FindStringSubmatch(ipStr)
	if ips == nil {
		return 0
	}
	temp, _ := strconv.Atoi(ips[1])
	c.port = uint16(temp)
	return c.port
}

func (c *Connection) GetConnectionId() uint32 {
	return c.cid
}

func (c *Connection) SetUid(uid string) {
	c.uid = uid
}

func (c *Connection) GetUid() string {
	return c.uid
}

func (c *Connection) Send(msg interface{}) bool {
	message := c.pro.Write(msg)
	_, _ = c.con.Write(message)
	return true
}

func (c *Connection) GetCon() net.Conn {
	return c.con
}

func (c *Connection) Close() {
	c.con.Close()
}

func (c *Connection) Read() (interface{}, error) {
	return c.pro.Read(c.con)
}

func NewConnection(con net.Conn, server network.ListenTcp, cid uint32) network.Connect {
	return &Connection{
		cid: cid,
		con: con,
		pro: server.GetProtocol(),
	}
}
