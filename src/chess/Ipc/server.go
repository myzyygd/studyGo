package Ipc

import (
	"encoding/json"
	"fmt"
	"log"
)

type Server interface {
	Name() string
	Handle(method string, param string) *Response
}

type Request struct {
	Method string `method`
	Params string `params`
}

type Response struct {
	Code string `code`
	Body string `body`
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c
			log.Println(request,"vvv")
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request params", err)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			log.Println("ddd")
			c <- string(b)
		}
		log.Println("Session closed")
	}(session)
	return session
}
