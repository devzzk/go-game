package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
