package states

import rl "github.com/gen2brain/raylib-go/raylib"

func UpdateWelcomeScreen(sm *StateMachine) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Title text settings
	title := "GONP!"
	fontSize := int32(100)
	// Center the title
	wMid, hMid := CenterText(title, fontSize)

	// Draw centered title
	rl.DrawText(title, wMid, hMid, fontSize, rl.White)

	// Draw other texts, using the same centering logic
	pressEnterText := "Press ENTER to start"
	pressEnterWidth, pressEnterHeight := CenterText(pressEnterText, 20)
	rl.DrawText(pressEnterText, pressEnterWidth, hMid+pressEnterHeight, 20, rl.Gray)

	pressEscapeText := "Press ESC to quit"
	pressEscapeWidth, pressEscapeHeight := CenterText(pressEscapeText, 20)
	rl.DrawText(pressEscapeText, pressEscapeWidth, hMid+pressEscapeHeight+50, 20, rl.Gray) // Slightly offset for separation

	rl.EndDrawing()

	if rl.IsKeyPressed(rl.KeyEnter) {
		sm.Set(StatePlaying)
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		sm.Set(StateGameOver)
	}
}

// CenterText returns the x and y coordinates to center text on the screen.
func CenterText(text string, fontSize int32) (int32, int32) {
	screenWidth := rl.GetScreenWidth()
	screenHeight := rl.GetScreenHeight()

	// Measure the width of the text
	textWidth := rl.MeasureText(text, fontSize)

	// Calculate centered position
	x := int32(screenWidth-int(textWidth)) / 2
	y := int32(screenHeight / 4) // Change this to control vertical positioning

	return x, y
}
