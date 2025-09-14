package systems

import (
	"github.com/TanZng/gonp/components"
	"github.com/andygeiss/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type inputSystem struct{}

func (a *inputSystem) Process(em ecs.EntityManager) int {
	player1 := em.Get("player1")
	player2 := em.Get("player2")

	if player1 != nil {
		handleInput(player1, ebiten.KeyW, ebiten.KeyS)
	}
	if player2 != nil {
		handleInput(player2, ebiten.KeyArrowUp, ebiten.KeyArrowDown)
	}

	return ecs.StateEngineContinue
}

func (a *inputSystem) Setup() {}

func (a *inputSystem) Teardown() {}

func NewInputSystem() ecs.System {
	return &inputSystem{}
}

func handleInput(entity *ecs.Entity, upKey, downKey ebiten.Key) {
	velocity := entity.Get(components.MaskVelocity).(*components.Velocity)

	// Reset vertical movement each frame
	velocity.Dy = 0

	if ebiten.IsKeyPressed(upKey) {
		velocity.Dy = -1 // move up
	} else if ebiten.IsKeyPressed(downKey) {
		velocity.Dy = 1 // move down
	}
}
