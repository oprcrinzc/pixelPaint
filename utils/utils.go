// Package utils for utilities function
package utils

import (
	"fmt"
	"image"
	"os"

	"pixelpaint/data"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/image/bmp"
)

func CheckIfCursorIsInArea(mousePos, origin rl.Vector2, width, height float32) bool {
	return (mousePos.X >= origin.X && mousePos.X <= origin.X+float32(width)) &&
		(mousePos.Y >= origin.Y && mousePos.Y <= origin.Y+float32(height))
}

func ExportToBmp() {
	img := image.NewRGBA(image.Rect(0, 0, data.SheetWidth, data.SheetWidth))
	for i := range data.SheetHeight {
		for j := range data.SheetWidth {
			img.Set(j, i, data.Sheet[i][j])
		}
	}

	if f, err := os.Create("output.bmp"); err == nil {
		defer f.Close()
		if err = bmp.Encode(f, img); err == nil {
			fmt.Println("export complete")
		} else {
			fmt.Println("export Fail:", err.Error())
		}
	} else {
		fmt.Println("export Fail:", err.Error())
	}
}

func CheckCursorHover() {
	if data.Hovers > 0 {
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
	data.Hovers = 0
}
