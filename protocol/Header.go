package protocol

import (
	"strings"
)

type Header struct {
	header map[string]string
}

func (*Header) Get(key string) string {
	panic("implement me")
}

func (h *Header) Set(data string) {
	h.header = map[string]string{}

	arr := strings.Split(data, "\r\n")
	for _, value := range arr[1:] {
		index := strings.Index(value, ":")
		if index >= 0 {
			h.header[value[:index]] = value[index+1:]
		}
	}
}
