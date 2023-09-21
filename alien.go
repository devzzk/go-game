package main

import (
	"log"
	"math/rand"

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
	LeftMove    bool
	HitNumber   int
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

	s := img.Bounds().Size()
	return &Alien{
		Image:       img,
		Width:       s.X,
		Height:      s.Y,
		X:           0,
		Y:           0,
		SpeedFactor: cfg.AlienSpeedFactor,
		LeftMove:    rand.Intn(100) > 50,
		HitNumber:   rand.Intn(3),
	}
}
