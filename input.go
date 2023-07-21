package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg string
}

func (i *Input) Update(ship *Ship, cfg *Config) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if ship.X > -float64(ship.Width)/2 {
			ship.X -= cfg.ShipSpeedFactor
		}

		i.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if ship.X < float64(cfg.ScreenWidth) - float64(ship.Width)/2 {
			ship.X += cfg.ShipSpeedFactor
		}

		i.msg = "right pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {

		i.msg = "space pressed"
	}
}
