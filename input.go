package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (g *Input) Player1() float32 {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		return -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		return 1
	}

	return float32(0.0)
}

func (g *Input) Player2() float32 {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		return -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		return 1
	}

	return float32(0.0)
}
