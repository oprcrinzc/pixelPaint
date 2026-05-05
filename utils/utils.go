// Package utils for utilities function
package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func CheckIfCursorIsInArea(mousePos, origin rl.Vector2, width, height float32) bool {
	return (mousePos.X >= origin.X && mousePos.X <= origin.X+float32(width)) &&
		(mousePos.Y >= origin.Y && mousePos.Y <= origin.Y+float32(height))
}
