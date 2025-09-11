package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y          float32
	Speed         float32
	Width, Height float32
	Paint         color.Color
	Score         float32
}

func NewPlayer(right bool, conf *Config) *Player {
	heightPaddle := float32(40.0)
	x, y := float32(conf.LayoutWidth-10), (float32(conf.LayoutHeight)-heightPaddle)/2
	if right {
		x, y = 10, (float32(conf.LayoutHeight)-heightPaddle)/2
	}
	return &Player{
		X:      x,
		Y:      y,
		Paint:  conf.Paint,
		Width:  4,
		Height: heightPaddle,
		Speed:  5,
	}
}

func (p *Player) Update() error {
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.X, p.Y, p.Width, p.Height, p.Paint, false)
}
