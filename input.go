package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (g *Input) Direction() float32 {
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		return -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		return 1
	}

	return float32(0.0)
}
