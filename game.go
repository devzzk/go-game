package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Game struct {
	Input   *Input
	ship    *Ship
	cfg     *Config
	bullets map[*Bullet]struct{}
	aliens  map[*Alien]struct{}
}

func (g *Game) Update() error {
	g.Input.Update(g)
	for bullet := range g.bullets {
		bullet.Y -= bullet.SpeedFactor
		if bullet.OutOfScreen() {
			delete(g.bullets, bullet)
		}
	}
	for alien := range g.aliens {
		alien.Y += alien.SpeedFactor
		if alien.X <= 0 {
			alien.LeftMove = false
		}
		if alien.X >= float64(g.cfg.ScreenWidth-alien.Width) {
			alien.LeftMove = true
		}

		if alien.LeftMove {
			alien.X -= alien.SpeedFactor
		} else {
			alien.X += alien.SpeedFactor
		}
	}

	g.CheckCollision()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)
	for bullet := range g.bullets {
		bullet.Draw(screen)
	}

	for alien := range g.aliens {
		alien.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

func (g *Game) AddBullet(bullet *Bullet) {
	if len(g.bullets) <= 10 {
		g.bullets[bullet] = struct{}{}
	}
}

func (g *Game) AddAlien(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

func (g *Game) CreateAliens() {
	alien := NewAlien(g.cfg)

	alien = NewAlien(g.cfg)
	alien.X = float64(rand.Intn(g.cfg.ScreenWidth - alien.Width))
	alien.Y = 0
	g.AddAlien(alien)
}

func (g Game) CheckCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if CheckCollision(bullet, alien) {
				alien.HitNumber -= 1
				if alien.HitNumber <= 0 {
					delete(g.aliens, alien)
				}
				delete(g.bullets, bullet)
			}
		}
	}
}

func CheckCollision(bullet *Bullet, alien *Alien) bool {
	alienTop, alienLeft := alien.Y, alien.X
	alienBottom, alienRight := alien.Y+float64(alien.Height), alien.X+float64(alien.Width)

	x, y := bullet.X, bullet.Y
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	// 右上角
	x, y = bullet.X+float64(bullet.Width), bullet.Y
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 左下角
	x, y = bullet.X, bullet.Y+float64(bullet.Height)
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 右下角
	x, y = bullet.X+float64(bullet.Width), bullet.Y+float64(bullet.Height)
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	return false
}

func NewGame() *Game {
	cfg := LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	g := &Game{
		Input: &Input{
			msg: "Hello world",
		},
		ship:    NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		cfg:     cfg,
		bullets: make(map[*Bullet]struct{}),
		aliens:  make(map[*Alien]struct{}),
	}
	return g
}
