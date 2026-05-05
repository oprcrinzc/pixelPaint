package main

import (
	"pixelpaint/data"
	"pixelpaint/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(600, 500, "Pixel Paint")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	// fmt.Println(string([]byte{'A', 'B'}))

	data.CurrentState = "CreateSheet"
	state.Load()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		state.Run()
		rl.EndDrawing()
	}
}
