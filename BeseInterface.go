package network

import (
	"net"
)

// tcp 服务端 or 客户端监听接口
type ListenTcp interface {
	SetUrl(address *Url)           // 设置监听地址
	Url() *Url                     // 地址
	SetEvent(event Event)          // 设置信息事件
	Event() Event                  // 信息事件
	SetProtocol(protocol Protocol) // 设置解析协议
	Protocol() Protocol            // 协议对象
	Close()                        // 主动关闭
	ListenAndServe()               // 启动监听，阻塞
	SetNewConnect(func(listen ListenTcp, conn net.Conn) Connect)
}

// 连接实例
type Connect interface {
	GetCon() net.Conn
	Close()
	Id() uint32
	SetUid(uid string)
	Uid() string
	Send(msg interface{}) bool
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
