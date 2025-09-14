package systems

import (
	"fmt"

	"github.com/TanZng/gonp/components"
	"github.com/andygeiss/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// renderingSystem ...
type renderingSystem struct {
	err           error
	title         string
	width, height int32
}

func NewRenderingSystem() *renderingSystem {
	return &renderingSystem{}
}

func (a *renderingSystem) Error() error {
	return a.err
}

func (a *renderingSystem) Setup() {
	rl.InitWindow(a.width, a.height, a.title)
}

func (a *renderingSystem) Process(em ecs.EntityManager) (state int) {
	// First check if app should stop.
	if rl.WindowShouldClose() {
		return ecs.StateEngineStop
	}
	// Clear the screen when the window is ready.
	if rl.IsWindowReady() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		a.renderEntities(em)
		a.renderScore(em)
		rl.EndDrawing()
	}
	return ecs.StateEngineContinue
}

func (a *renderingSystem) Teardown() {
	rl.CloseWindow()
}

func (a *renderingSystem) WithTitle(title string) *renderingSystem {
	a.title = title
	return a
}

func (a *renderingSystem) WithHeight(height float32) *renderingSystem {
	a.height = int32(height)
	return a
}

func (a *renderingSystem) WithWidth(width float32) *renderingSystem {
	a.width = int32(width)
	return a
}

func (a *renderingSystem) renderEntities(em ecs.EntityManager) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskSize) {
		position := e.Get(components.MaskPosition).(*components.Position)
		size := e.Get(components.MaskSize).(*components.Size)
		// Draw a bounding box
		rl.DrawRectangle(int32(position.X), int32(position.Y), int32(size.Width), int32(size.Height), size.Color)
	}
}

func (a *renderingSystem) renderScore(em ecs.EntityManager) {
	for _, e := range em.FilterByMask(components.MaskScore) {
		score := e.Get(components.MaskScore).(*components.Score)
		val := fmt.Sprintf("%d", score.Value)
		rl.DrawText(string(val), int32(score.X), int32(score.Y), 100, rl.Green)
	}
}
