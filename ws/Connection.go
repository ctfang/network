package ws

import (
	"bytes"
	"github.com/gorilla/websocket"
	"net"
	"network"
	"regexp"
	"strconv"
	"time"
)

type connection struct {
	cid    uint32
	uid    string
	con    *websocket.Conn
	pro    network.Protocol
	extend interface{}
	ip     uint32
	port   uint16

	write       chan []byte
	writeStatus bool
}

var MessageType = websocket.TextMessage

func NewConnection(con *websocket.Conn, server network.ListenTcp, cid uint32) network.Connect {
	conn := &connection{
		cid:   cid,
		con:   con,
		pro:   server.GetProtocol(),
		write: make(chan []byte, 245),
	}
	go conn.writeLoop()
	return conn
}

func (c *connection) writeLoop() {
	c.writeStatus = true
	for {
		select {
		case msg := <-c.write:
			_ = c.con.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.con.WriteMessage(MessageType, msg)

			if err != nil {
				c.Close()
				return
			}
		}
	}
}

func (c *connection) GetConnectionId() uint32 {
	return c.cid
}

func (c *connection) SetUid(uid string) {
	c.uid = uid
}

func (c *connection) GetUid() string {
	return c.uid
}

func (c *connection) Send(msg interface{}) bool {
	_ = c.con.SetWriteDeadline(time.Now().Add(writeWait))
	message := c.pro.Write(msg)
	c.write <- message

	return true
}

func (c *connection) GetCon() net.Conn {
	panic("websocket不能获取原连接")
}

func (c *connection) Close() {
	c.con.Close()
	if c.writeStatus {
		close(c.write)
		c.writeStatus = false
	}
}

func (c *connection) Read() (interface{}, error) {
	_, message, err := c.con.ReadMessage()
	if err != nil {
		return message, err
	}
	message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
	return message, nil
}

func (c *connection) SetExtend(extend interface{}) {
	c.extend = extend
}

func (c *connection) GetExtend() interface{} {
	return c.extend
}

func (c *connection) GetIp() uint32 {
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

func (c *connection) GetPort() uint16 {
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
