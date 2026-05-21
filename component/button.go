package component

import (
	"image/color"

	"pixelpaint/data"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// -------------------------------------------------------------------------

const (
	ButtonModeText = iota
	ButtonModeImg
)

// ---------------------------Button---------------------------------------------

type Button struct {
	origin     rl.Vector2
	width      int32
	height     int32
	color      color.RGBA
	hoverColor color.RGBA
	textColor  color.RGBA
	text       string
	img        rl.Texture2D
	mode       int
	fontSize   int32
	onClick    func()
}

func (b *Button) New(orin rl.Vector2) *Button {
	b.origin = orin
	b.onClick = func() {}
	b.img = rl.Texture2D{}
	b.mode = ButtonModeText
	return b
}

func (b *Button) SetText(text string) *Button {
	b.text = text
	b.fontSize = 30
	textSize := rl.MeasureText(text, b.fontSize)
	b.width = textSize + 10
	b.height = b.fontSize + 10
	b.color = color.RGBA{117, 90, 200, 255}
	b.hoverColor = color.RGBA{147, 120, 230, 255}
	b.textColor = color.RGBA{226, 191, 248, 255}
	return b
}

func (b *Button) SetImg(img rl.Texture2D) *Button {
	b.img = img
	b.width = img.Width
	b.height = img.Height
	return b
}

func (b *Button) SetWidthHeight(w, h int32) *Button {
	b.width = w
	b.height = h
	return b
}

func (b *Button) SetMode(mode int) *Button {
	b.mode = mode
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
		data.Hovers += 1
		switch b.mode {
		case ButtonModeText:
			rl.DrawRectangle(int32(b.origin.X), int32(b.origin.Y), b.width, b.height, b.hoverColor)
			rl.DrawText(b.text, int32(b.origin.X)+5, int32(b.origin.Y)+5, b.fontSize, b.textColor)
		case ButtonModeImg:
			rl.BeginBlendMode(rl.BlendAdditive)
			rl.DrawTextureEx(b.img, b.origin, 0, 1, rl.White)
			rl.EndBlendMode()
		}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			b.onClick()
		}

	} else {
		// rl.SetMouseCursor(rl.MouseCursorDefault)
		switch b.mode {
		case ButtonModeText:
			rl.DrawRectangle(int32(b.origin.X), int32(b.origin.Y), b.width, b.height, b.color)
			rl.DrawText(b.text, int32(b.origin.X)+5, int32(b.origin.Y)+5, b.fontSize, b.textColor)
		case ButtonModeImg:
			rl.DrawTextureEx(b.img, b.origin, 0, 1, rl.White)
		}
	}
}
