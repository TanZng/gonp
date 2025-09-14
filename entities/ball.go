package entities

import (
	"math/rand"
	"time"

	"github.com/TanZng/gonp/components"
)

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

func (b *Ball) Reset(x, y float32) {
	b.X = x
	b.Y = y
	b.Speed = 0.1
	b.Dx = -b.Dx
	b.Dy = randomValue()
}

func randomValue() float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	values := []float32{1.0, -1.0} // -1.0 has double the chance
	return values[r.Intn(len(values))]
}
