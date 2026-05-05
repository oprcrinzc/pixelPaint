package component

import (
	"image/color"
	"time"

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
	blinkTime   int64
	onBlink     int64
	offBlink    int64
	nextBlink   int64
	lastKeyNum  int32
}

// -------------------------------------------------------------------------

// New create InputText object
func (it *InputText) New(w int32, orin rl.Vector2) *InputText {
	it.width = w
	it.bgColor = color.RGBA{234, 135, 234, 255}
	it.textColor = color.RGBA{65, 36, 65, 255}
	it.cursorColor = color.RGBA{72, 34, 255, 255}
	it.focused = false
	it.value = ""
	it.origin = orin
	it.cursorPos = 0
	it.fontSize = 50
	it.height = it.fontSize
	it.blinkTime = 900
	it.onBlink = 900
	it.offBlink = 900
	it.nextBlink = time.Now().UnixMilli() + int64(it.blinkTime)
	it.lastKeyNum = 0
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

// SetFocus set focused to true
func (it *InputText) SetFocus(f bool) *InputText {
	it.focused = f
	return it
}

// -------------------------------------------------------------------------

// Render render InoutText
func (it *InputText) Render() {
	pdL := 10
	pdT := 5

	mousePos := rl.GetMousePosition()

	rl.DrawRectangle(int32(it.origin.X), int32(it.origin.Y), it.width, it.height, it.bgColor)
	rl.DrawText(it.value, int32(it.origin.X)+int32(pdL), int32(it.origin.Y)+int32(pdT), it.fontSize, it.textColor)

	if (mousePos.X >= it.origin.X && mousePos.X <= it.origin.X+float32(it.width)) &&
		(mousePos.Y >= it.origin.Y && mousePos.Y <= it.origin.Y+float32(it.height)) {
		rl.SetMouseCursor(rl.MouseCursorIBeam)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			it.SetFocus(true)
		}
	} else {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			it.SetFocus(false)
		}

		rl.SetMouseCursor(rl.MouseCursorDefault)
	}

	if it.focused {

		key := rl.GetKeyPressed()

		if (key >= 48 && key <= 57) || rl.IsKeyPressedRepeat(it.lastKeyNum) {
			it.nextBlink = it.nextBlink + 200
			if rl.IsKeyPressedRepeat(it.lastKeyNum) {
				it.value += string(byte(it.lastKeyNum))
			} else {
				it.value += string(byte(key))
				it.lastKeyNum = key
			}
		}

		lenValue := len(it.value)

		if (key == 259 || rl.IsKeyPressedRepeat(259)) && lenValue > 0 {
			it.nextBlink += 200
			it.value = it.value[:lenValue-1]
		}

		it.cursorPos = lenValue
		tsize := rl.MeasureText(it.value, it.fontSize)

		if time.Now().UnixMilli() <= it.nextBlink-it.offBlink {
			rl.DrawRectangle(int32(it.origin.X)+int32(pdL+int(tsize)+10),
				int32(it.origin.Y)+int32(pdT),
				int32(0.1*float64(it.fontSize)),
				it.fontSize-int32(2*(pdT)), it.cursorColor)
		}
		if time.Now().UnixMilli() > it.nextBlink {
			it.nextBlink = time.Now().UnixMilli() + int64(it.blinkTime) + it.onBlink
		}

	}
}

// -------------------------------------------------------------------------
