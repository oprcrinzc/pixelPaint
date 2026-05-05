package state

import (
	"strconv"

	"pixelpaint/data"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func StatePaintLoad(s *State) {
	PixelPrescaler := 20
	LinePrescaler := 1
	SheetPaintWidth := data.SheetWidth*PixelPrescaler + data.SheetWidth*LinePrescaler + 1
	SheetPaintHeight := data.SheetHeight*PixelPrescaler + data.SheetHeight*LinePrescaler + 1

	sheetTexture := rl.LoadRenderTexture(int32(SheetPaintWidth), int32(SheetPaintHeight))

	s.Storage["sheetRenderTexture"] = sheetTexture
	s.Storage["PixelPrescaler"] = PixelPrescaler
	s.Storage["LinePrescaler"] = LinePrescaler

	s.Storage["Scale"] = float32(1)

	rl.BeginTextureMode(sheetTexture)
	for i := range sheetTexture.Texture.Width {
		if i%int32(PixelPrescaler+1) == 0 {
			rl.DrawLine(i, 0, i+1, sheetTexture.Texture.Height, rl.Black)
		}
	}
	for i := range sheetTexture.Texture.Height {
		if i%int32(PixelPrescaler+1) == 0 {
			rl.DrawLine(0, i, sheetTexture.Texture.Width, i+1, rl.Black)
		}
	}
	rl.EndTextureMode()

	s.SetIsLoaded(true)
}
func StatePaintUnload(s *State) {}

// -------------------------------------------------------------------------

func StatePaintMain(s *State) {
	if !s.isLoaded {
		s.Load()
	}

	scale := float32(0)
	if s, ok := s.Storage["Scale"]; ok {
		scale = s.(float32)
	}

	if rl.IsKeyDown(rl.KeyMinus) {
		scale -= float32(rl.GetFrameTime() * 1)
	}
	if (rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)) && rl.IsKeyDown(rl.KeyEqual) {
		scale += float32(rl.GetFrameTime() * 1)
	}

	var sheet rl.RenderTexture2D

	if t, ok := s.Storage["sheetRenderTexture"]; ok {
		sheet = t.(rl.RenderTexture2D)
	}

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.DrawRectangle(0, 0, 100, data.HEIGHT, rl.DarkPurple)
	rl.DrawText("WxH:"+strconv.Itoa(data.SheetWidth)+"x"+strconv.Itoa(data.SheetHeight), 120, 10, 20, rl.Black)

	// rl.DrawTexture(sheet.Texture, 130, 30, rl.White)
	rl.DrawTextureEx(sheet.Texture, rl.NewVector2(130, 50), 0, float32(scale), rl.White)

	rl.EndDrawing()

	s.Storage["Scale"] = scale
}
