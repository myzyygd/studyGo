package Ipc

import (
	"testing"
	"fmt"
)

type EchoServer struct {
}

func (Server *EchoServer) Handle(method, params string) *Response {
	req := &Response{
		"测试",
		"测试中",
	}
	return req
}

func (Server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resq1, _ := client1.Call("From ", "client1")
	resq2, _ := client1.Call("From ", "client2")
	fmt.Println(resq1)
	fmt.Println(resq2)
	client1.Close()
	client2.Close()
}
