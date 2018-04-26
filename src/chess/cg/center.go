package cg

import (
	"sync"
	"chess/Ipc"
)

type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}

type Room struct{}

type CenterServer struct {
	Servers map[string]Ipc.IpcServer
	Players []*Player
	Rooms   []*Room
	Mutes   sync.RWMutex
}

var Mutes sync.RWMutex

func NewCenterServer() *CenterServer {
	Players := make([]*Player, 0)
	Servers := make(map[string]Ipc.IpcServer)
	Rooms := make([]*Room, 0)
	return &CenterServer{Servers, Players, Rooms, Mutes}
}
