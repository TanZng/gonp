package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (g *Input) Direction() float32 {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		fmt.Println("W")
		return -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		fmt.Println("S")
		return 1
	}

	return float32(0.0)
}
