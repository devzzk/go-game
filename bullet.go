package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Bullet struct {
	Image       *ebiten.Image
	Width       int
	Height      int
	X           float64
	Y           float64
	SpeedFactor float64
}

func (this *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(this.X, this.Y)
	screen.DrawImage(this.Image, op)
}

func (this *Bullet) OutOfScreen() bool {
	return this.Y < -float64(this.Height)
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		Image:       img,
		Width:       cfg.BulletWidth,
		Height:      cfg.BulletHeight,
		X:           ship.X + float64(ship.Width-cfg.BulletWidth)/2,
		Y:           float64(cfg.ScreenHeight - ship.Height - cfg.BulletHeight),
		SpeedFactor: cfg.BulletSpeedFactor,
	}
}
