package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	Image  *ebiten.Image
	Width  int
	Height int
	X      float64
	Y      float64
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(ship.X, ship.Y)

	screen.DrawImage(ship.Image, op)
}

func NewShip(screenWidth, screenHeight int) *Ship {

	img, _, err := ebitenutil.NewImageFromFile("./asset/ship.png")

	if err != nil {
		log.Fatal(err)
	}

	s := img.Bounds().Size()
	ship := &Ship{
		Image:  img,
		Width:  s.X,
		Height: s.Y,
		X:      float64(screenWidth-s.X) / 2,
		Y:      float64(screenHeight - s.Y),
	}
	return ship

}
