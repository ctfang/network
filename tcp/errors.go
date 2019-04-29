package tcp

import (
	"fmt"
	"github.com/ctfang/github.com/ctfang/network"
)

// 连接失败
type ListenError struct {
	address *network.Address
}

func (e *ListenError) Error() string {
	return fmt.Sprintf("连接失败 :%s", e.address.Str)
}
