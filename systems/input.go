package systems

import (
	"github.com/TanZng/gonp/components"
	"github.com/andygeiss/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type inputSystem struct{}

func (a *inputSystem) Process(em ecs.EntityManager) int {
	player1 := em.Get("player1").Get(components.MaskVelocity).(*components.Velocity)

	handleInput(player1, rl.KeyW, rl.KeyS)

	player2 := em.Get("player2").Get(components.MaskVelocity).(*components.Velocity)

	handleInput(player2, rl.KeyUp, rl.KeyDown)

	if rl.IsKeyDown(rl.KeyEscape) {
		return ecs.StateEngineStop
	}

	return ecs.StateEngineContinue
}

func (a *inputSystem) Setup() {}

func (a *inputSystem) Teardown() {}

func NewInputSystem() ecs.System {
	return &inputSystem{}
}

func handleInput(velocity *components.Velocity, upKey, downKey int32) {
	// Reset vertical movement each frame
	velocity.Dy = 0

	if rl.IsKeyDown(upKey) {
		velocity.Dy = -1 // move up
	} else if rl.IsKeyDown(downKey) {
		velocity.Dy = 1 // move down
	}
}
