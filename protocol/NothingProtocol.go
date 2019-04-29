package protocol

import (
	"network"
	"log"
	"net"
)

type NothingProtocol struct {
}

func (*NothingProtocol) Read(conn net.Conn) (interface{}, error) {
	panic("implement me")
}

func (*NothingProtocol) Write(msg interface{}) []byte {
	switch msg.(type) {
	case string:
		return []byte(msg.(string))
	case []byte:
		return msg.([]byte)
	default:
		log.Println("不认识的类型转换")
		return []byte("")
	}
}

func NewNothingProtocol() network.Protocol {
	return &NothingProtocol{}
}
