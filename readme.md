# TCP Server or Client 封装

> **简介:** 
对TCP服务端和客户端的封装，对于大多数应用，是不在意协议代码的实现的(因为协议本身只是个规范约束，于业务无关)，虽然go实现一个tcp很简单，但是要实现的话还是有不少沉余代码的。本库是对重复性代码进行封装。并且接口化。

使用本库只需要实现Event接口，协议部分
````
type Event interface {
	OnStart(listen ListenTcp)
	// 新链接
	OnConnect(connect Connect)
	// 新信息
	OnMessage(connect Connect, message []byte)
	// 链接关闭
	OnClose(connect Connect)
	// 发送错误
	OnError(listen ListenTcp, err error)
}
````

## 安装

````gotemplate
go get github.com/ctfang/network
````

## 使用

创建一个测试`LogicEvent`,`LogicEvent`必须实现 `Event` 接口

````go
package main

import (
	"github.com/ctfang/network"
	"log"
)

type LogicEvent struct {
}
func (*LogicEvent) OnStart(listen network.ListenTcp) {}
func (*LogicEvent) OnConnect(connect network.Connect) {}
func (*LogicEvent) OnMessage(connect network.Connect, message []byte) {
	connect.SendString("OnMessage")
}
func (*LogicEvent) OnClose(connect network.Connect) {}
func (*LogicEvent) OnError(listen network.ListenTcp, err error) {}
````

已经实现的协议有

<details open="open">
    <summary>ws、websocket 协议</summary>
    
````
func main() {
    server := network.NewServer("ws://127.0.0.1:8080")
    server.SetEvent(&LogicEvent{})
    server.ListenAndServe()
}
````
OR

````
func main() {
    server := network.NewClient("ws://127.0.0.1:8080")
    server.SetEvent(&LogicEvent{})
    server.ListenAndServe()
}
````
</details>


<details>
    <summary>text 协议</summary>
    
就是以回车为分隔的tcp协议，通常用来在命令行测试使用

````
func main() {
    server := network.NewServer("text://127.0.0.1:8081")
    server.SetEvent(&LogicEvent{})
    server.ListenAndServe()
}
````
OR

````
func main() {
    server := network.NewClient("text://127.0.0.1:8081")
    server.SetEvent(&LogicEvent{})
    server.ListenAndServe()
}
````
</details>


<details>
    <summary>pack 协议</summary>

自定义协议中常用的格式：包长(4位)+包文

````
func main() {
    server := network.NewServer("pack://127.0.0.1:8081")
    server.SetEvent(&LogicEvent{})
    server.ListenAndServe()
}
````
OR

````
func main() {
    server := network.NewClient("pack://127.0.0.1:8081")
    server.SetEvent(&LogicEvent{})
    server.ListenAndServe()
}
````
</details>

## 扩展

如果需要一些特别的协议格式，也可以创建 `YourProtocol` 协议解析文件进行自定义解析,实现接口 `Protocol`

````go
type Protocol interface {
    // 初始化
    Init()
    // 第一次连接，通常获取头信息
    OnConnect(conn net.Conn) (Header, error)
    // 读入处理
    Read(conn net.Conn) ([]byte, error)
    // 发送处理
    Write(conn net.Conn, msg []byte) error
}
````

````go
type YourProtocol struct {}
func (*YourProtocol) Init() { /* todo */ }
func (*YourProtocol) OnConnect(conn net.Conn) (Header, error) { /* todo */ }
func (*YourProtocol) Read(conn net.Conn) ([]byte, error) { /* todo */ }
func (*YourProtocol) Write(conn net.Conn, msg []byte) error  { /* todo */ }
````

使用自定义协议解析
````go

url := NewUrl(address)
server := Server{}
server.SetProtocol(&YourProtocol{})
server.SetUrl(url)
server.ListenAndServe()
````

## 示例

创建一个测试项目`test` 并且安装 `ctfang/network`
````
mkdir test
cd test
go mod ini test
go get github.com/ctfang/network
````
创建测试入口文件 main.go
创建测试入口文件 main.go
````go
package main

import (
	"github.com/ctfang/network"
	"log"
)

func main() {
    server := network.NewServer("ws://127.0.0.1:8080")
    server.SetEvent(&wsserverevent{})
    server.ListenAndServe()
}

type wsserverevent struct {
}

func (*wsserverevent) OnStart(listen network.ListenTcp) {

}

func (*wsserverevent) OnConnect(connect network.Connect) {
	connect.SendString("OnConnect")
}

func (*wsserverevent) OnMessage(connect network.Connect, message []byte) {
	log.Println(string(message))
	connect.SendString("OnMessage")
}

func (*wsserverevent) OnClose(connect network.Connect) {
	log.Println("OnClose")
}

func (*wsserverevent) OnError(listen network.ListenTcp, err error) {
	log.Println("OnError")
}

````