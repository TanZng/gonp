package entities

import "github.com/TanZng/gonp/components"

type Ball struct {
	*components.Position
	*components.Velocity
	*components.Size
}

func (b *Ball) IncreaseSpeed() {
	if b.Speed < 0.3 {
		b.Speed += 0.01
	}
}

func (b *Ball) ResetSpeed() {
	b.Speed = 0.1
}
