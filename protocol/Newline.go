package protocol

import (
	"bytes"
	"net"
)

/*
换行符基础协议
*/
type Newline struct {
	delim byte
}

func NewNewline() *Newline {
	return &Newline{
		delim: '\n',
	}
}

func (line *Newline) Read(conn net.Conn) (interface{}, error) {
	var message = make([]byte, 1024)
	var start = 0
	var count = 0
	for {
		w, err := conn.Read(message[start:])
		if err != nil {
			return "", err
		}
		index := bytes.IndexByte(message[start:start+w], line.delim)
		if index >= 0 {
			count = start + index
			break
		}
		start = start + w
	}

	return string(message[:count]), nil
}

func (line *Newline) Write(msg interface{}) []byte {
	switch msg.(type) {
	case []byte:
		return append(msg.([]byte), line.delim)
	case string:
		return append([]byte(msg.(string)), line.delim)
	}
	return []byte("")
}
