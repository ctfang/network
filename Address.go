package network

import (
	"net/url"
	"strconv"
	"strings"
)

type Url struct {
	Scheme string
	Host   string // host or host:port
	Path   string // path (relative paths may omit leading slash)
	Ip     string
	Port   uint16
}

func NewUrl(addr string) *Url {
	parse, err := url.Parse(addr)
	if err != nil {
		panic("地址格式错误")
	}
	arr := strings.Split(parse.Host, ":")
	port, _ := strconv.Atoi(parse.Port())
	return &Url{
		Scheme: parse.Scheme,
		Host:   parse.Host,
		Path:   parse.Path,
		Ip:     arr[0],
		Port:   uint16(port),
	}
}
