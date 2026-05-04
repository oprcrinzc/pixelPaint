package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(600, 500, "Pixel Paint")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Width", 100, 50, 50, rl.Black)
		rl.DrawText("Height", 100, 120, 50, rl.Black)
		rl.EndDrawing()
	}
}
