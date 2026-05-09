package main

import (
	"pixelpaint/data"
	"pixelpaint/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowResizable)
	rl.InitWindow(data.WIDTH, data.HEIGHT, "Pixel Paint")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(144)

	// fmt.Println(string([]byte{'A', 'B'}))
	//

	data.CurrentState = "CreateSheet"
	state.Load()

	for !rl.WindowShouldClose() {
		state.Run()
		rl.DrawFPS(500, 30)
	}
}
