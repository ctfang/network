package network

func NewClient(address string) ListenTcp {
	url := NewUrl(address)
	client := Client{}
	switch url.Scheme {
	case "ws":
		client.SetProtocol(&WsProtocol{})
	case "text":
		client.SetProtocol(&TextProtocol{})
	case "pack":
		client.SetProtocol(&PackageLenProtocol{})
	default:
		panic("ws or text")
	}
	client.SetUrl(url)
	return &client
}

func NewServer(address string) ListenTcp {
	url := NewUrl(address)

	server := Server{}
	switch url.Scheme {
	case "ws":
		server.SetProtocol(&WebsocketProtocol{})
	case "text":
		server.SetProtocol(&TextProtocol{})
	case "pack":
		server.SetProtocol(&PackageLenProtocol{})
	default:
		panic("ws、text、pack")
	}

	server.SetUrl(url)
	return &server
}
