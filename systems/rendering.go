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
	// rl.SetTargetFPS(120)
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

		rl.DrawText(fmt.Sprintln(rl.GetFPS(), "FPS"), a.width-100, a.height-30, 20, rl.Green)
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
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskSize | components.MaskColor) {
		position := e.Get(components.MaskPosition).(*components.Position)
		size := e.Get(components.MaskSize).(*components.Size)
		c := e.Get(components.MaskColor).(*components.Color)
		// Draw a bounding box
		rl.DrawRectangle(int32(position.X), int32(position.Y), int32(size.Width), int32(size.Height), c.Value)
	}
}

func (a *renderingSystem) renderScore(em ecs.EntityManager) {
	for _, e := range em.FilterByMask(components.MaskScore | components.MaskColor) {
		score := e.Get(components.MaskScore).(*components.Score)
		c := e.Get(components.MaskColor).(*components.Color)
		val := fmt.Sprintf("%d", score.Value)
		rl.DrawText(string(val), int32(score.X), int32(score.Y), 100, c.Value)
	}
}
