package network

import (
	"fmt"
)

// 连接失败
type ListenError struct {
	address *Url
}

func (e *ListenError) Error() string {
	return fmt.Sprintf("连接失败 :%s", e.address.Host)
}
