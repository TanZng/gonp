package main

import (
	"github.com/TanZng/gonp/components"
	"github.com/TanZng/gonp/systems"
	"github.com/andygeiss/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var width, height float32 = 800, 600
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("topwall", []ecs.Component{
		components.NewPosition().WithX(1).WithY(1),
		components.NewSize(width, 10, rl.Green),
	}))
	em.Add(ecs.NewEntity("downwall", []ecs.Component{
		components.NewPosition().WithX(1).WithY(height - 10),
		components.NewSize(width, 10, rl.Green),
	}))
	em.Add(ecs.NewEntity("middle", []ecs.Component{
		components.NewPosition().WithX(width / 2).WithY(0),
		components.NewSize(10, height, rl.Green),
	}))
	em.Add(ecs.NewEntity("ball", []ecs.Component{
		components.NewPosition().WithX(width / 2).WithY(height / 2),
		components.NewSize(10, 10, rl.Green),
		components.NewVelocity().WithDx(1).WithDy(1).WithSpeed(0.1),
	}))
	em.Add(ecs.NewEntity("player1", []ecs.Component{
		components.NewPosition().WithX(20).WithY(height / 2),
		components.NewSize(10, 100, rl.Green),
		components.NewVelocity().WithSpeed(0.3),
		components.NewScore(width/4, 20),
	}))
	em.Add(ecs.NewEntity("player2", []ecs.Component{
		components.NewPosition().WithX(width - 28).WithY(height / 2),
		components.NewSize(10, 100, rl.Green),
		components.NewVelocity().WithSpeed(0.3),
		components.NewScore(width/4*3, 20),
	}))
	sm := ecs.NewSystemManager()
	sm.Add(
		systems.NewInputSystem(),
		systems.NewMovementSystem(height),
		systems.NewCollisionSystem().WithWidth(width).WithHeight(height),
		systems.NewRenderingSystem().WithWidth(width).WithHeight(height),
	)
	de := ecs.NewDefaultEngine(em, sm)
	de.Setup()
	defer de.Teardown()
	de.Run()
}
