package systems

import (
	"github.com/TanZng/gonp/components"
	"github.com/andygeiss/ecs"
)

type collisionSystem struct {
	width, height int32
}

func (a *collisionSystem) Process(em ecs.EntityManager) (state int) {
	ball := em.Get("ball")
	ballPos := ball.Get(components.MaskPosition).(*components.Position)
	ballVel := ball.Get(components.MaskVelocity).(*components.Velocity)

	// Check ball out-of-bounds (left/right edges)
	if ballPos.X < 0 || ballPos.X > float32(a.width) {
		// Reset ball position, update scores etc.
		ballPos.X = float32(a.width) / 2
		ballPos.Y = float32(a.height) / 2
		// You might want to reset velocity here too
	}

	// Bounce off top/bottom walls
	if ballPos.Y < 0 || ballPos.Y > float32(a.height) {
		ballVel.Dy = -ballVel.Dy
	}

	// Check collisions with players
	ballSize := ball.Get(components.MaskSize).(*components.Size)
	playerEntities := em.FilterByNames("player")
	for _, player := range playerEntities {
		playerPos := player.Get(components.MaskPosition).(*components.Position)
		playerSize := player.Get(components.MaskSize).(*components.Size)

		if ballPos.X < playerPos.X+playerSize.Width && ballPos.X+ballSize.Height > playerPos.X &&
			ballPos.Y < playerPos.Y+playerSize.Height && ballPos.Y+ballSize.Height > playerPos.Y {
			ballVel.Dx = -ballVel.Dx
			ballVel.IncreaseSpeed()
		}
	}

	return ecs.StateEngineContinue
}

func (a *collisionSystem) Setup(em ecs.EntityManager) {
}

func (a *collisionSystem) Teardown() {}

func NewCollisionSystem() *collisionSystem {
	return &collisionSystem{}
}
