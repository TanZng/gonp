package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	conf *Config

	ball    *Ball
	players []*Player

	world *World

	input *Input

	ms *MovementSystem
}

func (g *Game) Update() error {
	g.ms.Player(g.players[0], g.world, g.input)
	g.ms.Ball(g.ball, g.world, g.players)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	g.ball.Draw(screen)
	g.players[0].Draw(screen)
	g.players[1].Draw(screen)
	g.world.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.conf.LayoutWidth, g.conf.LayoutHeight
}

func main() {
	conf := NewConfig()
	ebiten.SetWindowSize(conf.ScreenWidth, conf.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	ball := NewBall(conf)
	player := NewPlayer(true, conf)
	enemy := NewPlayer(false, conf)
	world := NewWorld(conf)
	input := NewInput()

	if err := ebiten.RunGame(&Game{conf, ball, []*Player{player, enemy}, world, input, &MovementSystem{}}); err != nil {
		log.Fatal(err)
	}
}
