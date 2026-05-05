package state

import (
	"fmt"
	"math"
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
	rl.SetTextureFilter(sheetTexture.Texture, rl.FilterBilinear)

	s.Storage["sheetRenderTexture"] = sheetTexture
	s.Storage["PixelPrescaler"] = PixelPrescaler
	s.Storage["LinePrescaler"] = LinePrescaler

	s.Storage["Scale"] = float32(1)

	drawGridFunc := func() {
		rl.BeginTextureMode(sheetTexture)
		rl.ClearBackground(rl.White)
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
	}

	drawGridFunc()

	s.Storage["DrawGrid"] = drawGridFunc

	s.Storage["CanvasOrigin"] = rl.NewVector2(130, 50)

	s.SetIsLoaded(true)
}
func StatePaintUnload(s *State) {}

// -------------------------------------------------------------------------

func StatePaintMain(s *State) {
	if !s.isLoaded {
		s.Load()
	}

	mousePos := rl.GetMousePosition()

	// -------------------------------------------------------------------------

	scale := float32(0)
	if s, ok := s.Storage["Scale"]; ok {
		scale = s.(float32)
	}

	PixelPrescaler := 0
	if pp, ok := s.Storage["PixelPrescaler"]; ok {
		PixelPrescaler = pp.(int)
	}

	var CanvasOrigin rl.Vector2
	if a, ok := s.Storage["CanvasOrigin"]; ok {
		CanvasOrigin = a.(rl.Vector2)
	}

	// ---------------------------[ Calculate X, Y in canvas ]---------------------------------------------

	deltaX := mousePos.X - CanvasOrigin.X
	deltaY := mousePos.Y - CanvasOrigin.Y
	x := math.Floor(float64(deltaX) / (float64(PixelPrescaler+1) * float64(scale)))
	y := math.Floor(float64(deltaY) / (float64(PixelPrescaler+1) * float64(scale)))

	rl.DrawText(fmt.Sprintf("mouse X,Y: (%d,%d), scale: %f", int(x), int(y), scale), 300, 10, 20, rl.Black)

	// -------------------------------------------------------------------------

	// utils.CheckIfCursorIsInArea(mousePos, CanvasOrigin, width float32, height float32)

	if rl.IsKeyDown(rl.KeyMinus) {
		scale -= float32(rl.GetFrameTime() * 0.5)
		if f, ok := s.Storage["DrawGrid"]; ok {
			f.(func())()
		}
	}
	if (rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)) && rl.IsKeyDown(rl.KeyEqual) {
		scale += float32(rl.GetFrameTime() * 0.5)
		if f, ok := s.Storage["DrawGrid"]; ok {
			f.(func())()
		}
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
