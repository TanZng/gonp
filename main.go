package main

import (
	"github.com/TanZng/gonp/components"
	"github.com/TanZng/gonp/states"
	"github.com/TanZng/gonp/systems"
	"github.com/andygeiss/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  float32 = 800
	screenHeight float32 = 600

	wallThickness = 5
	wallXOffset   = 1
	wallYOffset   = 1

	paddleWidth  = 5
	paddleHeight = 100
	paddleSpeed  = 0.3

	player1XOffset = 20
	player2XOffset = 30

	scoreYOffset = 20

	ballSize  = 10
	ballSpeed = 0.1
	ballDx    = 1
	ballDy    = 1
)

var defaultColor = rl.White

func main() {
	sm := states.NewStateMachine(states.StateWelcome)

	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Pong with ECS State Machine")

	defer rl.CloseWindow()
	var engine = setupEcsEngine()

	runGame := true

	for runGame {
		switch {
		case sm.Is(states.StateWelcome):
			states.UpdateWelcomeScreen(sm)

		case sm.Is(states.StatePlaying):
			engine.Run()
			sm.Set(states.StatePause)

		case sm.Is(states.StatePause):
			states.UpdatePauseScreen(sm)

		case sm.Is(states.StateGameOver):
			runGame = false
		}
	}
}

// Entity Helpers

func newWall(name string, x, y, w, h float32) *ecs.Entity {
	return ecs.NewEntity(name, []ecs.Component{
		components.NewPosition().WithX(x).WithY(y),
		components.NewSize(w, h),
		components.NewColor(defaultColor),
	})
}

func newBall(x, y float32) *ecs.Entity {
	return ecs.NewEntity("ball", []ecs.Component{
		components.NewPosition().WithX(x).WithY(y),
		components.NewSize(ballSize, ballSize),
		components.NewVelocity().WithDx(ballDx).WithDy(ballDy).WithSpeed(ballSpeed),
		components.NewColor(defaultColor),
	})
}

func newPlayer(name string, x, y, scoreX, scoreY float32) *ecs.Entity {
	return ecs.NewEntity(name, []ecs.Component{
		components.NewPosition().WithX(x).WithY(y),
		components.NewSize(paddleWidth, paddleHeight),
		components.NewVelocity().WithSpeed(paddleSpeed),
		components.NewScore(scoreX, scoreY),
		components.NewColor(defaultColor),
	})
}

func setupEcsEngine() ecs.Engine {
	em := ecs.NewEntityManager()

	// Walls
	em.Add(newWall("topwall", wallXOffset, wallYOffset, screenWidth, wallThickness))
	em.Add(newWall("downwall", wallXOffset, screenHeight-wallThickness, screenWidth, wallThickness))
	em.Add(newWall("middle", screenWidth/2, 0, wallThickness, screenHeight))

	// Ball
	em.Add(newBall(screenWidth/2, screenHeight/2))

	// Players
	em.Add(newPlayer("player1", player1XOffset, screenHeight/2, screenWidth/4, scoreYOffset))
	em.Add(newPlayer("player2", screenWidth-player2XOffset, screenHeight/2, screenWidth*3/4, scoreYOffset))

	// Systems
	sm := ecs.NewSystemManager()
	sm.Add(
		systems.NewInputSystem(),
		systems.NewMovementSystem(screenHeight),
		systems.NewCollisionSystem().WithWidth(screenWidth).WithHeight(screenHeight),
		systems.NewRenderingSystem().WithWidth(screenWidth).WithHeight(screenHeight),
	)

	return ecs.NewDefaultEngine(em, sm)
}
