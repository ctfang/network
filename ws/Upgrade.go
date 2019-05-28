package ws

import (
	"github.com/ctfang/network"
)

type Upgrade struct {
	GET                 string
	Upgrade             string
	Connection          string
	Host                string
	Origin              string
	SecWebSocketKey     string
	SecWebSocketVersion string
}

func (u *Upgrade) String() string {
	str := "GET " + u.GET + "\r\n"
	str += "Upgrade: " + u.Upgrade + "\r\n"
	str += "Connection: " + u.Connection + "\r\n"
	str += "Host: " + u.Host + "\r\n"
	str += "Origin: " + u.Origin + "\r\n"
	str += "Sec-WebSocket-Key: " + u.SecWebSocketKey + "\r\n"
	str += "Sec-WebSocket-Version: " + u.SecWebSocketVersion + "\r\n\r\n"

	return str
}

func NewUpgrade(url *network.Url) *Upgrade {
	return &Upgrade{
		GET:                 "/ HTTP/1.1",
		Upgrade:             "websocket",
		Connection:          "Upgrade",
		Host:                url.Host,
		Origin:              "http://" + url.Host,
		SecWebSocketKey:     "sN9cRrP/n9NdMgdcy2VJFQ==",
		SecWebSocketVersion: "13",
	}
}
