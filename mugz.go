package main

import (
	"fmt"

	"github.com/makyo/mugz/game"
	"github.com/makyo/mugz/server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("We had some bad times...", r)
		}
		fmt.Println("Goodbye!")
	}()

	g, err := game.LoadGame("game.json")
	if err != nil {
		panic(err)
	}

	server, err := server.New(g, ":5523")
	if err != nil {
		panic(err)
	}

	if err = server.Listen(); err != nil {
		panic(err)
	}

}
