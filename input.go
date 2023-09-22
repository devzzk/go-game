package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Input struct {
	msg                 string
	lastBulletTime      time.Time
	lastCreateAllenTime time.Time
}

func (this *Input) Update(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.ship.X > -float64(g.ship.Width)/2 {
			g.ship.X -= g.cfg.ShipSpeedFactor
		}

		this.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.ship.X < float64(g.cfg.ScreenWidth)-float64(g.ship.Width)/2 {
			g.ship.X += g.cfg.ShipSpeedFactor
		}

		this.msg = "right pressed"
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if g.ship.X > -float64(g.ship.Height)/2 {
			g.ship.Y -= g.cfg.ShipSpeedFactor
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.ship.Y < float64(g.cfg.ScreenHeight)-float64(g.ship.Height)/2 {
			g.ship.Y += g.cfg.ShipSpeedFactor
		}
	}

	if time.Now().Sub(this.lastCreateAllenTime).Milliseconds() > g.cfg.CreateAllenInterval {
		g.CreateAliens()
		this.lastCreateAllenTime = time.Now()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// int 无法和 int64 比较 类型不同无法比较
		if time.Now().Sub(this.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
			bullet := NewBullet(g.cfg, g.ship)
			g.AddBullet(bullet)
			this.lastBulletTime = time.Now()
		}
	}
}
