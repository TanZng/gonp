package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	X, Y  float32
	Size  float32
	Paint color.Color
}

func NewBall(conf *Config) *Ball {
	return &Ball{
		X:     float32(conf.LayoutWidth / 2),
		Y:     float32(conf.LayoutHeight / 2),
		Size:  5,
		Paint: conf.Paint,
	}
}

func (b *Ball) Update() {

}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.X, b.Y, b.Size, b.Paint, false)
}
