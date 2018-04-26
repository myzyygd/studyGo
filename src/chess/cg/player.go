package cg

import (
	"fmt"
)

type Player struct {
	Name  string        `name`
	Level int           `level`
	Exp   int           `exp`
	Room  int           `room`
	mq    chan *Message `message`
}


func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}
	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "Receipd message", msg.Content)
		}
	}(player)
	return player
}
