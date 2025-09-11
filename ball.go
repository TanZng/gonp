package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	initialSpeed = 1
)

type Ball struct {
	X, Y   float32
	DirX   float32 // dir horizontal, -1 o 1
	DirY   float32 // dir vertical, -1 o 1
	Radius float32
	Speed  float32
	Paint  color.Color
}

func NewBall(conf *Config) *Ball {
	return &Ball{
		X:      float32(conf.LayoutWidth / 2),
		Y:      float32(conf.LayoutHeight / 2),
		DirX:   1, // mueve a la derecha
		DirY:   1, // mueve hacia abajo
		Speed:  initialSpeed,
		Radius: 4,
		Paint:  conf.Paint,
	}
}

func (b *Ball) Update() {
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Radius, b.Radius, b.Paint, false)
}

func (b *Ball) IncreaseSpeed() {
	if b.Speed < 8 {
		b.Speed += 0.5
	}
}

func (b *Ball) Respawn(x, y float32) {
	b.X, b.Y = x, y
	b.DirX = b.DirX * -1
	b.DirY = randomValue()
	b.Speed = initialSpeed
}

func randomValue() float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	values := []float32{1.0, -1.0} // -1.0 has double the chance
	return values[r.Intn(len(values))]
}
