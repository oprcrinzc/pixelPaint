// Package data is central of this program
package data

import "image/color"

const (
	WIDTH  int32 = 1280
	HEIGHT int32 = 720
)

var (
	CurrentState string
	Sheet        map[int]map[int]color.RGBA = make(map[int]map[int]color.RGBA)

	SheetWidth  int = 0
	SheetHeight int = 0

	// IsHover bool = false
	Hovers uint = 0
)
