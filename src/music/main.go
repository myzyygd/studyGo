package main

import (
	"fmt"
	"music/library"
	"bufio"
	"os"
	"strings"
	"music/mp"
	"log"
)

var lib *library.MusicManger
var Id int = 1

func main() {
	fmt.Println(`
	Enter following commands to control the player:
	lib list --View the existing music lib
	lib add <name><artist><genre><source><type> -- Add a music to the music lib
	lib remove <name> --Remove the specified music from the lib
	play <name> -- Play the specified music
     `)

	lib = library.NewMusicManger()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter Command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		log.Println(line)
		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			HandleLibCommand(tokens)
		} else if tokens[1] == "play" {
			HandlePlayCommand(tokens)
		} else {
			fmt.Println("Unsupport Command", tokens[0])
		}
	}

}

func HandleLibCommand(tokens []string) {
	switch tokens[1] {
	case "list":
		if lib.Len() <= 0 {
			fmt.Println("音乐库没有东西")
			break
		}
		for i := 0; i < lib.Len(); i++ {
			m, _ := lib.Get(i)
			fmt.Println(i+1, ":", m.Name, m.ArtList, m.Source, m.Type)
		}
	case "add":
		if len(tokens) != 7 {
			fmt.Println("USAGE : lib add <name><artist><genre><source><type> (7 argv)")
			return
		}
		Id++
		m := &library.MusicEntry{
			Id:      Id,
			Name:    tokens[2],
			ArtList: tokens[3],
			Genre:   tokens[4],
			Source:  tokens[5],
			Type:    tokens[6],
		}
		lib.Add(m)
	case "remove":
		if len(tokens) != 2 {
			fmt.Println("lib remove <name> --Remove the specified music from the lib")
			return
		}
		lib.RemoveByName(tokens[2])
	default:
		fmt.Println("Unrecogized lib command: ", tokens[1])

	}
}

func HandlePlayCommand(tokens []string) {
	fmt.Println(tokens)
	if len(tokens) != 2 {
		fmt.Println("USAGE : play <name>")
		return
	}
	m := lib.Find(tokens[1])
	if m == nil {
		fmt.Println("The Music Not Found")
		return
	}
	mp.Play(m.Source, m.Type)
}
