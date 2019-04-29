package network

import (
	"strconv"
	"strings"
)

type Address struct {
	// 监听地址，只包含域名端口
	Str  string
	Ip   string
	Port uint16
	// 地址，不包含域名端口
	Path string
}

func NewAddress(addr string) *Address {
	strS := strings.Split(addr, ":")
	if len(strS) != 2 {
		panic("格式错误")
	}
	Port, _ := strconv.ParseInt(strS[1], 10, 64)
	index := strings.Index(addr, "/")
	var pathStr string
	var Str string
	if index > 1 {
		pathStr = addr[index:]
		Str = addr[:index]
	} else {
		pathStr = "/"
		Str = addr
	}
	return &Address{
		Str:  Str,
		Ip:   strS[0],
		Port: uint16(Port),
		Path: pathStr,
	}
}
