package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/ctfang/network"
	"net"
)

type WebsocketProtocol struct {
	// 本地缓冲区
	cacheByte []byte
	// 缓冲长度
	cacheCount int
}

/*
GET / HTTP/1.1
Upgrade: websocket
Connection: Upgrade
Host: example.com
Origin: http://example.com
Sec-WebSocket-Key: sN9cRrP/n9NdMgdcy2VJFQ==
Sec-WebSocket-Version: 13
*/
func (w *WebsocketProtocol) AsClient(conn net.Conn) (network.Header, error) {
	w.cacheByte = make([]byte, 0)
	// 发送请求头
	strHeader := "GET / HTTP/1.1\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nHost: "
	strHeader += conn.RemoteAddr().String() + "\r\n"
	strHeader += "Origin: http://" + conn.RemoteAddr().String() + "\r\n"
	strHeader += "Sec-WebSocket-Key:sN9cRrP/n9NdMgdcy2VJFQ==\r\n"
	strHeader += "Sec-WebSocket-Version: 13\r\n\r\n"
	conn.Write([]byte(strHeader))

	// 获取协议头
	byteHeader, err := w.getHeader(conn)
	header := Header{}
	header.Set(string(byteHeader))
	return &header, err
}

func (*WebsocketProtocol) AsServer(conn net.Conn) (network.Header, error) {
	panic("implement me")
}

func (w *WebsocketProtocol) Read(conn net.Conn) ([]byte, error) {
	// 第一个字节：FIN + RSV1-3 + OPCODE
	opcodeByte, err := w.readConnOrCache(conn, 1)
	FIN := opcodeByte[0] >> 7
	// RSV1 := opcodeByte[0] >> 6 & 1 // 自定义协议
	// RSV2 := opcodeByte[0] >> 5 & 1 // 自定义协议
	// RSV3 := opcodeByte[0] >> 4 & 1 // 自定义协议
	// OPCODE := opcodeByte[0] & 15
	//
	// log.Println(RSV1, RSV2, RSV3, OPCODE)

	// 第二个字节，Mask + Payload length
	payloadLenByte, err := w.readConnOrCache(conn, 1)
	payloadLen := int(payloadLenByte[0] & 0x7F) // 有效负载
	mask := payloadLenByte[0] >> 7

	switch payloadLen {
	case 126: // 两个字节表示的是一个16进制无符号数，这个数用来表示传输数据的长度
		temLen, _ := w.readConnOrCache(conn, 2)
		payloadLen = int(binary.BigEndian.Uint16(temLen))
	case 127: // 8个字节表示的一个64位无符合数，这个数用来表示传输数据的长度
		temLen, _ := w.readConnOrCache(conn, 4)
		payloadLen = int(binary.BigEndian.Uint64(temLen))
	}

	msg := make([]byte, payloadLen)
	if mask == 0 {
		// 服务端信息，没有掩码
		msg, _ = w.readConnOrCache(conn, payloadLen)
	} else {
		tem, _ := w.readConnOrCache(conn, payloadLen)
		// 掩码读取
		maskingByte, _ := w.readConnOrCache(conn, 4)
		for i := 0; i < payloadLen; i++ {
			msg[i] = tem[i+4] ^ maskingByte[i%4]
		}
	}

	if FIN == 1 {
		// 最后的消息片断
		return msg, err
	}

	nextMsg, err := w.Read(conn)
	msg = append(msg, nextMsg...)
	return msg, err
}

func (*WebsocketProtocol) Write(msg []byte) []byte {
	panic("implement me")
}

// 读取指定长度数据
func (w *WebsocketProtocol) readConnOrCache(conn net.Conn, count int) ([]byte, error) {
	if w.cacheCount > 0 {
		// 拥有缓冲数据
		if count <= w.cacheCount {
			// 缓冲数据比需要的还要大，直接拿取
			msg := w.cacheByte[:count]
			w.cacheCount = w.cacheCount - count
			w.cacheByte = w.cacheByte[count:]
			return msg, nil
		} else {
			// 缓冲数据不足，剩余需要的位数，多读取一点，可以优化速度
			data := make([]byte, count+512)
			cacheCount, err := conn.Read(data)
			if err != nil {
				return []byte(""), errors.New("读取数据失败")
			}
			w.cacheCount = w.cacheCount + cacheCount
			w.cacheByte = append(w.cacheByte, data...)
			return w.readConnOrCache(conn, count)
		}
	} else {
		// 缓冲是空的
		data := make([]byte, 1024)
		cacheCount, err := conn.Read(data)
		if err != nil {
			return []byte(""), errors.New("读取数据失败")
		}
		w.cacheCount = cacheCount
		w.cacheByte = append(w.cacheByte, data...)
		return w.readConnOrCache(conn, count)
	}
}

func (w *WebsocketProtocol) getHeader(conn net.Conn) ([]byte, error) {
	data := make([]byte, 1024)
	count, err := conn.Read(data)
	if err != nil {
		return []byte(""), errors.New("获取协议头信息错误")
	}

	sep := []byte("\r\n\r\n")
	index := bytes.Index(data, sep)
	if index > 0 {
		// 一般都是比较小数据的头部
		if count != index+4 {
			w.cacheByte = append(w.cacheByte, data[index+4:count]...)
			w.cacheCount = count - index - 4
		}
		return data[:index], nil
	} else {
		return []byte(""), errors.New("协议头异常，不存在分隔符也不应超过1024位")
	}
}
