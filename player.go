package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y          float32
	Width, Height float32
	Paint         color.Color
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
	}
}

func (b *Player) Update() error {
	return nil
}

func (b *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Width, b.Height, b.Paint, false)
}
