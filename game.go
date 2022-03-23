package main

import "fmt"

type Game struct {
	Room string
}

func (g *Game) StartGame(room string) {
	g.Room = room
}

func (g *Game) PrintGameState() {
	fmt.Println("Room name: " + g.Room)
}
