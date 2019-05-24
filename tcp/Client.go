package tcp

import (
	"github.com/ctfang/network/ws"
	"log"
	"net"
)

type Client struct {
	ListenTcp
}

func (c *Client) ListenAndServe() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", c.url.Host)
	c.conn, err = net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		go c.event.OnError(c, &ListenError{c.url})
		log.Printf("tcp client 启动失败, err : %v\n", err.Error())
		return
	}
	c.id += 1
	go c.event.OnStart(c)

	if c.newConnect == nil {
		switch c.Url().Scheme {
		case "ws":
			c.newConnect = ws.NewConnect
		default:
			c.newConnect = NewConnect
		}
	}
	Connect := c.newConnect(c, c.conn)
	go c.event.OnConnect(Connect)
	for {
		data := make([]byte, 1024)
		count, _ := c.conn.Read(data)
		log.Println(count)
	}
}

//
// func (client c *Client) Url() *network.Url {
// 	return client.url
// }
// func (client c *Client) Conn() net.Conn {
// 	return client.conn
// }
// func (client c *Client) SetEvent(event network.Event) {
// 	client.event = event
// }
//
// func (client c *Client) Event() network.Event {
// 	return client.event
// }
//
// func (client c *Client) SetProtocol(protocol network.Protocol) {
// 	client.protocol = protocol
// }
//
// // 主动关闭连接
// func (client c *Client) Close() {
// 	_ = client.con.Close()
// }
//
// func (client c *Client) ListenAndServe() {
// 	tcpAddr, err := net.ResolveTCPAddr("tcp4", client.url.Str)
// 	conn, err := net.DialTCP("tcp", nil, tcpAddr)
//
//
// 	if err != nil {
// 		go client.event.OnError(client.event, &ListenError{client.url})
// 		log.Printf("tcp client 启动失败, err : %v\n", err.Error())
// 		return
// 	}
// 	client.conn = conn
// 	client.id += 1
// 	go client.event.OnStart(client)
// 	client.newConnection(conn)
// }
//
// /*
// 新的连接
// */
// func (client c *Client) newConnection(con net.Conn) {
// 	var connection = NewConnection(con, client, client.lastId)
// 	event := client.GetConnectionEvent()
// 	go event.OnConnect(connection)
// 	defer event.OnClose(connection)
//
// 	for {
// 		message, err := connection.Read()
// 		if err != nil {
// 			con.Close()
// 			break
// 		}
// 		go event.OnMessage(connection, message)
// 	}
// }
