package component

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	origin     rl.Vector2
	width      int32
	height     int32
	color      color.RGBA
	hoverColor color.RGBA
	textColor  color.RGBA
	text       string
	fontSize   int32
	onClick    func()
}

func (b *Button) New(t string, orin rl.Vector2) *Button {
	b.origin = orin
	b.text = t
	b.fontSize = 30
	textSize := rl.MeasureText(t, b.fontSize)
	b.width = textSize + 10
	b.height = b.fontSize + 10
	b.color = color.RGBA{117, 90, 200, 255}
	b.hoverColor = color.RGBA{147, 120, 230, 255}
	b.textColor = color.RGBA{226, 191, 248, 255}
	b.onClick = func() {}
	return b
}

func (b *Button) Bind(f func()) *Button {
	b.onClick = f
	return b
}

func (b *Button) Render() {
	mousePos := rl.GetMousePosition()

	if (mousePos.X >= b.origin.X && mousePos.X <= b.origin.X+float32(b.width)) &&
		(mousePos.Y >= b.origin.Y && mousePos.Y <= b.origin.Y+float32(b.height)) {
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
		rl.DrawRectangle(int32(b.origin.X), int32(b.origin.Y), b.width, b.height, b.hoverColor)
		rl.DrawText(b.text, int32(b.origin.X)+5, int32(b.origin.Y)+5, b.fontSize, b.textColor)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			b.onClick()
		}
	} else {
		rl.DrawRectangle(int32(b.origin.X), int32(b.origin.Y), b.width, b.height, b.color)
		rl.DrawText(b.text, int32(b.origin.X)+5, int32(b.origin.Y)+5, b.fontSize, b.textColor)
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
}
