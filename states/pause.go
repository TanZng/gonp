package states

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpdatePauseScreen(sm *StateMachine) {
	rl.BeginDrawing()
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.NewColor(0, 0, 0, 2))

	// Center the "Paused" text
	title := "Paused"
	fontSize := int32(40)
	wMid, hMid := CenterText(title, fontSize)
	rl.DrawText(title, wMid, hMid, fontSize, rl.White)

	// Center the other text options
	pressEnterText := "Press ENTER to Resume"
	pressEnterWidth, pressEnterHeight := CenterText(pressEnterText, 20)
	rl.DrawText(pressEnterText, pressEnterWidth, hMid+pressEnterHeight, 20, rl.White)

	pressMText := "Press M to Main Menu"
	pressMWidth, pressMHeight := CenterText(pressMText, 20)
	rl.DrawText(pressMText, pressMWidth, hMid+pressMHeight+40, 20, rl.White) // Slightly offset for separation

	pressEscapeText := "Press ESC to Quit"
	pressEscapeWidth, pressEscapeHeight := CenterText(pressEscapeText, 20)
	rl.DrawText(pressEscapeText, pressEscapeWidth, hMid+pressEscapeHeight+80, 20, rl.White) // Slightly offset for separation

	rl.EndDrawing()

	// Handle input for pause menu
	if rl.IsKeyPressed(rl.KeyEnter) {
		sm.Set(StatePlaying) // Resume game
	}

	if rl.IsKeyPressed(rl.KeyM) {
		sm.Set(StateWelcome) // Return to main menu
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		sm.Set(StateGameOver) // Quit game
	}
}
