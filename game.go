package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input *Input
	ship  *Ship
	cfg   *Config
}

func (g *Game) Update() error {
	g.input.Update(g.ship, g.cfg)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

func NewGame() *Game {
	cfg := LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	return &Game{
		input: &Input{
			msg: "Hello world",
		},
		ship: NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		cfg:  cfg,
	}
}
