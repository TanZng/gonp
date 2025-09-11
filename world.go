package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type World struct {
	Width, Height float32
	Paint         color.Color

	strokeWidth  float32
	startScreen  float32
	middleScreen float32
}

func NewWorld(conf *Config) *World {
	return &World{
		Width:  float32(conf.LayoutWidth),
		Height: float32(conf.LayoutHeight),
		Paint:  conf.Paint,

		strokeWidth:  1,
		startScreen:  0,
		middleScreen: float32(conf.LayoutWidth / 2),
	}
}

func (w *World) Update() error {
	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	// Top Line -----
	vector.StrokeLine(screen, w.startScreen, 1, w.Width, 1, w.strokeWidth, w.Paint, false)
	// Bot Line ____
	vector.StrokeLine(screen, w.startScreen, w.Height, w.Width, w.Height, w.strokeWidth, w.Paint, false)
	// Mid Line |
	vector.StrokeLine(screen, w.middleScreen, 1, w.middleScreen, w.Height, w.strokeWidth, w.Paint, false)
}
