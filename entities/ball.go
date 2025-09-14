package entities

import (
	"github.com/TanZng/gonp/components"
	"github.com/TanZng/gonp/utils"
)

const (
	maxSpeed      = 0.3
	increaseSpeed = 0.01
)

type Ball struct {
	*components.Position
	*components.Velocity
	*components.Size
}

// IncreaseSpeed increases the speed up to a defined maximum.
func (b *Ball) IncreaseSpeed() {
	if b.Speed < maxSpeed {
		b.Speed += increaseSpeed
	}
}

// Reset repositions the ball and sets new direction/speed.
func (b *Ball) Reset(x, y float32) {
	b.X = x
	b.Y = y
	b.Speed = 0.1
	b.Dx = -b.Dx
	b.Dy = utils.RandomValueBetween(1.0, -1.0)
}
