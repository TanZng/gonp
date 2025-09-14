package entities

import "github.com/TanZng/gonp/components"

type Player struct {
	*components.Position
	*components.Velocity
	*components.Size
}
