package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Input struct {
	msg            string
	lastBulletTime time.Time
}

func (i *Input) Update(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.ship.X > -float64(g.ship.Width)/2 {
			g.ship.X -= g.cfg.ShipSpeedFactor
		}

		i.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.ship.X < float64(g.cfg.ScreenWidth)-float64(g.ship.Width)/2 {
			g.ship.X += g.cfg.ShipSpeedFactor
		}

		i.msg = "right pressed"
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// int 无法和 int64 比较 类型不同无法比较
		if time.Now().Sub(i.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
			bullet := NewBullet(g.cfg, g.ship)
			g.AddBullet(bullet)
			i.lastBulletTime = time.Now()
		}
	}
}
