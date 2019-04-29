package network

import (
	"net"
)

// tcp 服务端 or 客户端监听接口
type ListenTcp interface {
	SetAddress(address *Address)
	GetAddress() *Address
	SetConnectionEvent(event Event)
	GetConnectionEvent() Event
	SetProtocol(protocol Protocol)
	GetProtocol() Protocol
	ListenAndServe()
}

// 连接实例
type Connect interface {
	GetCon() net.Conn
	Close()
	GetConnectionId() uint32
	SetUid(uid string)
	GetUid() string
	Send(msg interface{}) bool
	Read() (interface{}, error)
	SetExtend(extend interface{})
	GetExtend() interface{}
	GetIp() uint32
	GetPort() uint16
}

type Event interface {
	OnStart(listen ListenTcp)
	// 新链接
	OnConnect(connect Connect)
	// 新信息
	OnMessage(connect Connect, message interface{})
	// 链接关闭
	OnClose(connect Connect)
	// 发送错误
	OnError(listen ListenTcp, err error)
}

type Protocol interface {
	// 读入处理
	Read(conn net.Conn) (interface{}, error)
	// 发送处理
	Write(msg interface{}) []byte
}
