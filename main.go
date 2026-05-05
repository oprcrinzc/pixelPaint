package main

import (
	"pixelpaint/component"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(600, 500, "Pixel Paint")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	it1 := component.InputText{}
	it1.New(300, rl.NewVector2(300, 50))

	// fmt.Println(string([]byte{'A', 'B'}))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Width", 100, 50, 50, rl.Black)
		rl.DrawText("Height", 100, 120, 50, rl.Black)
		it1.Render()
		rl.EndDrawing()
	}
}
