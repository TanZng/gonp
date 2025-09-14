package systems

import (
	"math/rand"
	"time"

	"github.com/TanZng/gonp/components"
	"github.com/TanZng/gonp/entities"
	"github.com/andygeiss/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type collisionSystem struct {
	width, height int32
}

func (a *collisionSystem) Process(em ecs.EntityManager) (state int) {
	ball := entities.Ball{
		Position: em.Get("ball").Get(components.MaskPosition).(*components.Position),
		Velocity: em.Get("ball").Get(components.MaskVelocity).(*components.Velocity),
		Size:     em.Get("ball").Get(components.MaskSize).(*components.Size),
	}

	// Check ball out-of-bounds (left/right edges)
	if ball.X < 0 || ball.X > float32(a.width) {
		// Reset ball position, update scores etc.
		ball.X = float32(a.width) / 2
		ball.Y = float32(a.height) / 2
		// You might want to reset velocity here too
		ball.Dx = -ball.Dx
		ball.Dy = randomValue()
		ball.ResetSpeed()
	}

	// Bounce off top/bottom walls
	if ball.Y < 0 || ball.Y > float32(a.height) {
		ball.Dy = -ball.Dy
	}

	// Check collisions with players
	player1 := em.Get("player1")
	player2 := em.Get("player2")
	playerEntities := []*ecs.Entity{player1, player2}
	for _, player := range playerEntities {
		player := entities.Player{
			Position: player.Get(components.MaskPosition).(*components.Position),
			Size:     player.Get(components.MaskSize).(*components.Size),
		}

		if checkCollision(ball, player) {
			ball.Dx = -ball.Dx
			ball.IncreaseSpeed()
		}
	}

	return ecs.StateEngineContinue
}

func (a *collisionSystem) Setup() {
}

func (a *collisionSystem) Teardown() {}

func (a *collisionSystem) WithHeight(height float32) *collisionSystem {
	a.height = int32(height)
	return a
}

func (a *collisionSystem) WithWidth(width float32) *collisionSystem {
	a.width = int32(width)
	return a
}

func NewCollisionSystem() *collisionSystem {
	return &collisionSystem{}
}

func randomValue() float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	values := []float32{1.0, -1.0} // -1.0 has double the chance
	return values[r.Intn(len(values))]
}

// checkCollision between ball and paddle, and determine side
func checkCollision(ball entities.Ball, player entities.Player) bool {
	closestX := rl.Clamp(ball.X, player.X, player.X+player.Width)
	closestY := rl.Clamp(ball.Y, player.Y, player.Y+player.Height)

	dx := ball.X - closestX
	dy := ball.Y - closestY

	distanceSq := dx*dx + dy*dy
	radiusSq := ball.Width * ball.Height

	if distanceSq <= radiusSq {
		// Determine where it hit
		if closestX == player.X || closestX == player.X+player.Width {
			return true
		}
		if closestY == player.Y || closestY == player.Y+player.Height {
			return false
		}
		return true // fallback
	}

	return false
}
