package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	renderWidth  = 540
	renderHeight = 960
	renderScale  = 1
	screenWidth  = 540
	screenHeight = 960
	targetFPS    = 30
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "playground")

	rl.SetTargetFPS(targetFPS)

	vb := NewVerticalBounce()
	hb := NewHorizontalBounce()
	db := NewDiagonalBounce()
	dnb := NewDiagonalNormalBounce()

	for !rl.WindowShouldClose() {

		rl.ClearBackground(rl.Black)

		rl.BeginDrawing()

		if rl.IsKeyPressed(rl.KeyPageUp) {
			vb.SpeedUp()
			hb.SpeedUp()
			db.SpeedUp()
			dnb.SpeedUp()
		}
		if rl.IsKeyPressed(rl.KeyPageDown) {
			vb.SpeedDown()
			hb.SpeedDown()
			db.SpeedDown()
			dnb.SpeedDown()
		}

		vb.Update()
		hb.Update()
		db.Update()
		dnb.Update()

		vb.Draw()
		hb.Draw()
		db.Draw()
		dnb.Draw()

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
