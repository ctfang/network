package tcp

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/protocol"
	"log"
	"net"
)

type Client struct {
	address  *network.Address
	event    network.Event
	protocol network.Protocol
	con      net.Conn
	lastId   uint32
}

func NewClient() network.ListenTcp {
	return &Client{}
}

func (client *Client) SetAddress(address *network.Address) {
	client.address = address
}

func (client *Client) GetAddress() *network.Address {
	return client.address
}

func (client *Client) SetConnectionEvent(event network.Event) {
	client.event = event
}

func (client *Client) GetConnectionEvent() network.Event {
	return client.event
}

func (client *Client) SetProtocol(protocol network.Protocol) {
	client.protocol = protocol
}

func (client *Client) GetProtocol() network.Protocol {
	if client.protocol == nil {
		client.SetProtocol(protocol.NewNewline())
	}
	return client.protocol
}

// 主动关闭连接
func (client *Client) Close() {
	_ = client.con.Close()
}

func (client *Client) ListenAndServe() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", client.address.Str)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	event := client.GetConnectionEvent()

	if err != nil {
		go event.OnError(client, &ListenError{client.address})
		log.Printf("tcp client 启动失败, err : %v\n", err.Error())
		return
	}
	client.con = conn
	client.lastId += 1
	go event.OnStart(client)
	client.newConnection(conn)
}

/*
新的连接
*/
func (client *Client) newConnection(con net.Conn) {
	var connection = NewConnection(con, client, client.lastId)
	event := client.GetConnectionEvent()
	go event.OnConnect(connection)
	defer event.OnClose(connection)

	for {
		message, err := connection.Read()
		if err != nil {
			con.Close()
			break
		}
		go event.OnMessage(connection, message)
	}
}
