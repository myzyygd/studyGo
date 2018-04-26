package Ipc

import (
	"encoding/json"
	"log"
	"runtime"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method string, params string) (resp *Response, err error) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	req := &Request{method, params}
	var resp1 Response
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	log.Println(string(b), "aa")
	client.conn <- string(b)
	str := <-client.conn
	log.Println("cccc")
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
