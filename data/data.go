// Package data is central of this program
package data

import "image/color"

var (
	CurrentState string
	Sheet        map[int]map[int]color.RGBA = make(map[int]map[int]color.RGBA)

	SheetWidth  int = 0
	SheetHeight int = 0
)
