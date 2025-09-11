package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	normalFontSize = 24
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
	g.ms.Enemy(g.players[1], g.world, g.input)
	g.ms.Ball(g.ball, g.world, g.players)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	g.ball.Draw(screen)
	g.players[0].Draw(screen)
	g.players[1].Draw(screen)
	g.world.Draw(screen)

	// Draw the text
	half := float64((g.world.Width) / 4)

	drawScore(screen, half, g.players[0])
	drawScore(screen, half*3, g.players[1])
}

func drawScore(screen *ebiten.Image, position float64, player *Player) {
	fontCenter := float64(normalFontSize / 2)

	op := &text.DrawOptions{}
	op.GeoM.Translate(position-fontCenter, 10)
	op.ColorScale.ScaleWithColor(player.Paint)
	sampleText := fmt.Sprintf("%v", player.Score)
	text.Draw(screen, sampleText, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, op)
}

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.conf.LayoutWidth, g.conf.LayoutHeight
}

func main() {
	conf := NewConfig()
	ebiten.SetWindowSize(conf.ScreenWidth, conf.ScreenHeight)
	ebiten.SetWindowTitle("GONP!")

	ball := NewBall(conf)
	player := NewPlayer(true, conf)
	enemy := NewPlayer(false, conf)
	world := NewWorld(conf)
	input := NewInput()

	if err := ebiten.RunGame(&Game{conf, ball, []*Player{player, enemy}, world, input, &MovementSystem{}}); err != nil {
		log.Fatal(err)
	}
}
