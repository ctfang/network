package ws

import (
	"github.com/ctfang/network"
	"github.com/ctfang/network/protocol"
	"github.com/ctfang/network/tcp"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Server struct {
	upgrader websocket.Upgrader
	address  *network.Address
	event    network.Event
	protocol network.Protocol
	listener *websocket.Conn
	lastId   uint32
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func NewServer() network.ListenTcp {
	return &Server{}
}

func (server *Server) SetAddress(address *network.Address) {
	server.address = address
}

func (server *Server) GetAddress() *network.Address {
	return server.address
}

func (server *Server) SetConnectionEvent(event network.Event) {
	server.event = event
}

func (server *Server) GetConnectionEvent() network.Event {
	if server.event == nil {
		panic("没有设置事件")
	}
	return server.event
}

func (server *Server) SetProtocol(protocol network.Protocol) {
	server.protocol = protocol
}

func (server *Server) GetProtocol() network.Protocol {
	if server.protocol == nil {
		server.protocol = protocol.NewNothingProtocol()
	}
	return server.protocol
}

func (server *Server) ListenAndServe() {
	server.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	event := server.GetConnectionEvent()
	address := server.GetAddress()

	go event.OnStart(server)

	http.HandleFunc(address.Path, server.Upgrade)
	err := http.ListenAndServe(address.Str, nil)
	if err != nil {
		event.OnError(server, &tcp.ListenError{})
		log.Fatal("websocket 启动失败: ", err)
	}
}
func (server *Server) Close() {
	_ = server.listener.Close()
}

func (server *Server) Upgrade(w http.ResponseWriter, r *http.Request) {
	con, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	server.listener = con
	// 信息size上限
	con.SetReadLimit(maxMessageSize)
	// 设置底层网络连接的读取截止日期。读取超时后，websocket连接状态已损坏，所有将来的读取都将返回错误。t的零值意味着读取不会超时。
	// con.SetReadDeadline(time.Now().Add(pongWait))
	// Pong 信息
	con.SetPongHandler(func(string) error { _ = con.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	server.lastId++
	var connection = NewConnection(con, server, server.lastId)
	go server.event.OnConnect(connection)
	go server.readPump(con, connection)
}

func (server *Server) readPump(con *websocket.Conn, connection network.Connect) {
	defer server.event.OnClose(connection)

	for {
		message, err := connection.Read()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("websocket error: %v", err)
			}
			con.Close()
			break
		}
		go server.event.OnMessage(connection, message)
	}
}
