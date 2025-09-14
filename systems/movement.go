package systems

import (
	"github.com/TanZng/gonp/components"
	"github.com/andygeiss/ecs"
)

type movementSystem struct {
	height float32
}

func (a *movementSystem) Process(em ecs.EntityManager) (state int) {
	ball := em.Get("ball")
	ballPos := ball.Get(components.MaskPosition).(*components.Position)
	ballVel := ball.Get(components.MaskVelocity).(*components.Velocity)
	ballPos.X += ballVel.Dx * ballVel.Speed
	ballPos.Y += ballVel.Dy * ballVel.Speed

	// Calculate the next position of the players.
	player1 := em.Get("player1")
	player2 := em.Get("player2")

	playerEntities := []*ecs.Entity{player1, player2}
	for _, player := range playerEntities {
		pos := player.Get(components.MaskPosition).(*components.Position)
		vel := player.Get(components.MaskVelocity).(*components.Velocity)
		size := player.Get(components.MaskSize).(*components.Size)

		vy := vel.Dy * vel.Speed
		nextY := pos.Y + vy
		if nextY >= 0 && nextY <= a.height-size.Height {
			pos.Y += vy
		}
	}

	return ecs.StateEngineContinue
}

func (a *movementSystem) Setup() {}

func (a *movementSystem) Teardown() {}

func NewMovementSystem(height float32) ecs.System {
	return &movementSystem{height}
}
