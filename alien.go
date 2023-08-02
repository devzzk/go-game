package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Alien struct {
	Image       *ebiten.Image
	Width       int
	Height      int
	X           float64
	Y           float64
	SpeedFactor float64
}

func (this Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(this.X, this.Y)
	screen.DrawImage(this.Image, op)
}

func NewAlien(cfg *Config) *Alien {
	img, _, err := ebitenutil.NewImageFromFile("./asset/alien.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	return &Alien{
		Image:       img,
		Width:       width,
		Height:      height,
		X:           0,
		Y:           0,
		SpeedFactor: cfg.AlienSpeedFactor,
	}
}
