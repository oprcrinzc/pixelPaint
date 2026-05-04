package component

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// -------------------------------------------------------------------------

type InputText struct {
	width       int32
	height      int32
	bgColor     color.RGBA
	textColor   color.RGBA
	cursorColor color.RGBA
	focused     bool
	value       string
	origin      rl.Vector2
	cursorPos   int
	fontSize    int32
}

// -------------------------------------------------------------------------

// New create InputText object
func (it *InputText) New(w, h int32, orin rl.Vector2) *InputText {
	it.width = w
	it.height = h
	it.bgColor = color.RGBA{234, 135, 234, 255}
	it.textColor = color.RGBA{65, 36, 65, 255}
	it.cursorColor = color.RGBA{72, 34, 255, 255}
	it.focused = false
	it.value = ""
	it.origin = orin
	it.cursorPos = 0
	it.fontSize = 50
	return it
}

// -------------------------------------------------------------------------
// Get

// GetValue read value from InputText
func (it *InputText) GetValue() string {
	return it.value
}

// IsFocused read focused from InputText
func (it *InputText) IsFocused() bool {
	return it.focused
}

// -------------------------------------------------------------------------
// Set

// Focus set focused to true
func (it *InputText) Focus() *InputText {
	it.focused = true
	return it
}

// -------------------------------------------------------------------------

// Render render InoutText
func (it *InputText) Render() {
	rl.DrawRectangle(int32(it.origin.X), int32(it.origin.Y), it.width, it.height, it.bgColor)
	if it.focused {
		rl.DrawRectangle(int32(20+it.cursorPos+10/2), 10, int32(0.25*float64(it.fontSize)), it.fontSize, it.cursorColor)
	}
}

// -------------------------------------------------------------------------
