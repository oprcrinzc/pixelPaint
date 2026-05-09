package state

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"pixelpaint/component"
	"pixelpaint/data"
	"pixelpaint/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func StatePaintLoad(s *State) {
	PixelPrescaler := 14
	LinePrescaler := 1
	SheetPaintWidth := data.SheetWidth*PixelPrescaler + data.SheetWidth*LinePrescaler + 1
	SheetPaintHeight := data.SheetHeight*PixelPrescaler + data.SheetHeight*LinePrescaler + 1

	sheetTexture := rl.LoadRenderTexture(int32(SheetPaintWidth), int32(SheetPaintHeight))
	rl.SetTextureFilter(sheetTexture.Texture, rl.FilterBilinear)
	rl.SetTextureFilter(sheetTexture.Texture, rl.FilterPoint)

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

	updatePixel := func() {
		rl.BeginTextureMode(sheetTexture)
		py := 0
		px := 0
		for i := range sheetTexture.Texture.Height {
			if i%int32(PixelPrescaler+1) == 1 {
				px = 0
				for j := range sheetTexture.Texture.Width {
					if j%int32(PixelPrescaler+1) == 1 {
						if col, ok := data.Sheet[int(py)][int(px)]; ok {
							rl.DrawRectangle(j, i, int32(PixelPrescaler), int32(PixelPrescaler), col)
						}
						px += 1
					}
				}
				py += 1
			}
		}
		// fmt.Println("px, py:", px, py)

		rl.EndTextureMode()
	}

	drawGridFunc()
	updatePixel()

	s.Storage["DrawGrid"] = drawGridFunc
	s.Storage["UpdatePixel"] = updatePixel

	// ---------------------------Center Canvas ---------------------------------------------

	screenHeight := rl.GetScreenHeight()
	screenWidth := rl.GetScreenWidth()

	canvasOriginX := float32(int(screenWidth)/2 - int(sheetTexture.Texture.Width)/2)
	canvasOriginY := float32(int(screenHeight)/2 - int(sheetTexture.Texture.Height)/2)

	s.Storage["CanvasOrigin"] = rl.NewVector2(-canvasOriginX, -canvasOriginY)
	s.Storage["CanvasOriginDefault"] = rl.NewVector2(-canvasOriginX, -canvasOriginY)

	// -------------------------------------------------------------------------

	s.Storage["CurrentMode"] = "draw"

	drawModeImg := rl.LoadTexture("./krita/export/drawModeImg.png")
	s.Storage["DrawModeImg"] = drawModeImg
	drawModeBtn := component.Button{}
	drawModeBtn.New(rl.NewVector2(20, 100)).
		SetMode(component.ButtonModeImg).
		SetImg(drawModeImg).
		Bind(func() {
			s.Storage["CurrentMode"] = "draw"
		})
	s.Storage["DrawModeBtn"] = drawModeBtn

	eraseModeImg := rl.LoadTexture("./krita/export/eraseModeImg.png")
	s.Storage["EraseModeImg"] = eraseModeImg
	eraseModeBtn := component.Button{}
	eraseModeBtn.New(rl.NewVector2(20, 200)).
		SetMode(component.ButtonModeImg).
		SetImg(eraseModeImg).
		Bind(func() {
			s.Storage["CurrentMode"] = "erase"
		})
	s.Storage["EraseModeBtn"] = eraseModeBtn

	exportImg := rl.LoadTexture("./krita/export/export.png")
	s.Storage["ExportImg"] = exportImg
	exportBtn := component.Button{}
	exportBtn.New(rl.NewVector2(20, 300)).
		SetMode(component.ButtonModeImg).
		SetImg(exportImg).
		Bind(func() {
			s.Storage["CurrentMode"] = "none"
			utils.ExportToBmp()
		})
	s.Storage["ExportBtn"] = exportBtn

	// -----------------------------Test ColorSheer------------------------------------------

	cs := component.ColorSheet{}
	cs.New(256, 256).Bind(func(c *component.ColorSheet) {
		Bs := uint8(0)

		if bs, ok := s.Storage["TEST_COLORSHEET_B_INPUT"]; ok {
			Bs = bs.(uint8)
		}
		for i := range c.Height {
			for j := range c.Width {
				c.Set(i, j, color.RGBA{uint8(i), uint8(j), uint8(Bs), 255})
			}
		}
	})
	s.Storage["TEST_COLORSHEET_TEXTURE"] = cs.Render()
	s.Storage["TEST_COLORSHEET_RENDER_FUNC"] = cs.Render

	// -------------------------------------------------------------------------
	s.SetIsLoaded(true)
}

func StatePaintUnload(s *State) {
	if i, ok := s.Storage["DrawModeImg"]; ok {
		ii := i.(rl.Texture2D)
		rl.UnloadTexture(ii)
	}
	if i, ok := s.Storage["EraseModeImg"]; ok {
		ii := i.(rl.Texture2D)
		rl.UnloadTexture(ii)
	}
	if i, ok := s.Storage["ExportImg"]; ok {
		ii := i.(rl.Texture2D)
		rl.UnloadTexture(ii)
	}
}

// -------------------------------------------------------------------------

func StatePaintMain(s *State) {
	if !s.isLoaded {
		s.Load()
	}

	// -------------------------------------------------------------------------

	mousePos := rl.GetMousePosition()
	// screenWidth := rl.GetScreenWidth()
	screenHeight := rl.GetScreenHeight()

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

	var CanvasOriginDefault rl.Vector2
	if a, ok := s.Storage["CanvasOriginDefault"]; ok {
		CanvasOriginDefault = a.(rl.Vector2)
	}

	var sheet rl.RenderTexture2D
	if t, ok := s.Storage["sheetRenderTexture"]; ok {
		sheet = t.(rl.RenderTexture2D)
	}

	// ---------------------------[ Calculate X, Y in canvas ]---------------------------------------------

	deltaX := mousePos.X - CanvasOrigin.X*(-1)
	deltaY := mousePos.Y - CanvasOrigin.Y*(-1)
	x := math.Floor(float64(deltaX) / (float64(PixelPrescaler+1) * float64(scale)))
	y := math.Floor(float64(deltaY) / (float64(PixelPrescaler+1) * float64(scale)))

	rl.DrawText(fmt.Sprintf("mouse X,Y: (%d,%d), scale: %f", int(x), int(y), scale), 300, 10, 20, rl.Black)

	// -------------------------------------------------------------------------

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		if x >= 0 && y >= 0 && x < float64(data.SheetWidth) && y < float64(data.SheetHeight) {
			c := color.RGBA{0, 0, 0, 255}
			if currentMode_, ok := s.Storage["CurrentMode"]; ok {
				currentMode := currentMode_.(string)
				switch currentMode {
				case "draw":
					c = color.RGBA{0, 0, 0, 255}
				case "erase":
					c = color.RGBA{255, 255, 255, 255}
				}
			}

			data.Sheet[int(y)][int(x)] = c
			// fmt.Println(data.Sheet)
			if f, ok := s.Storage["UpdatePixel"]; ok {
				f.(func())()
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyTab) {
		if f, ok := s.Storage["TEST_COLORSHEET_RENDER_FUNC"]; ok {
			s.Storage["TEST_COLORSHEET_TEXTURE"] = f.(func() rl.Texture2D)()
		}
	}

	// utils.CheckIfCursorIsInArea(mousePos, CanvasOrigin, width float32, height float32)

	if rl.IsKeyDown(rl.KeyMinus) {
		scale -= float32(rl.GetFrameTime() * 0.5)
		CanvasOrigin.X = CanvasOriginDefault.X + float32(float32(sheet.Texture.Width)*scale-float32(sheet.Texture.Width))/2
		CanvasOrigin.Y = CanvasOriginDefault.Y + float32(float32(sheet.Texture.Height)*scale-float32(sheet.Texture.Height))/2
		if f, ok := s.Storage["DrawGrid"]; ok {
			f.(func())()
		}
		if f, ok := s.Storage["UpdatePixel"]; ok {
			f.(func())()
		}
	}
	if (rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)) && rl.IsKeyDown(rl.KeyEqual) {
		scale += float32(rl.GetFrameTime() * 0.5)
		CanvasOrigin.X = CanvasOriginDefault.X + float32(float32(sheet.Texture.Width)*scale-float32(sheet.Texture.Width))/2
		CanvasOrigin.Y = CanvasOriginDefault.Y + float32(float32(sheet.Texture.Height)*scale-float32(sheet.Texture.Height))/2
		if f, ok := s.Storage["DrawGrid"]; ok {
			f.(func())()
		}
		if f, ok := s.Storage["UpdatePixel"]; ok {
			f.(func())()
		}
	}

	moveSpeed := 100
	if rl.IsKeyDown(rl.KeyLeft) {
		CanvasOrigin.X += rl.GetFrameTime() * float32(moveSpeed)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		CanvasOrigin.X -= rl.GetFrameTime() * float32(moveSpeed)
	}
	if rl.IsKeyDown(rl.KeyUp) {
		CanvasOrigin.Y += rl.GetFrameTime() * float32(moveSpeed)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		CanvasOrigin.Y -= rl.GetFrameTime() * float32(moveSpeed)
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	// ----------------------------Draw tools bar-------------------------------------------

	rl.DrawRectangle(0, 0, 100, int32(screenHeight), rl.DarkPurple)
	// fmt.Println(screenWidth)

	if d, ok := s.Storage["DrawModeBtn"]; ok {
		dbtn := d.(component.Button)
		dbtn.Render()
	}

	if d, ok := s.Storage["EraseModeBtn"]; ok {
		dbtn := d.(component.Button)
		dbtn.Render()
	}

	if d, ok := s.Storage["ExportBtn"]; ok {
		dbtn := d.(component.Button)
		dbtn.Render()
	}

	// -------------------------------------------------------------------------

	rl.DrawText("WxH:"+strconv.Itoa(data.SheetWidth)+"x"+strconv.Itoa(data.SheetHeight), 120, 10, 20, rl.Black)

	// rl.DrawTexture(sheet.Texture, 130, 30, rl.White)
	//	rl.DrawTextureEx(sheet.Texture, CanvasOrigin, 0, float32(scale), rl.White)

	rl.DrawTexturePro(sheet.Texture,
		rl.NewRectangle(0, 0, float32(sheet.Texture.Width), -float32(sheet.Texture.Height)),
		rl.NewRectangle(0, 0, float32(sheet.Texture.Width)*scale, float32(sheet.Texture.Height)*scale),
		CanvasOrigin,
		0,
		rl.White)

	rl.DrawText(fmt.Sprintf("Origin: %v", CanvasOrigin), 120, 40, 20, rl.Black)

	if c, ok := s.Storage["TEST_COLORSHEET_TEXTURE"]; ok {
		cst := c.(rl.Texture2D)
		rl.DrawTexture(cst, 90, 90, rl.White)
	}

	rl.EndDrawing()

	s.Storage["Scale"] = scale
	s.Storage["CanvasOrigin"] = CanvasOrigin

	s.Storage["TEST_COLORSHEET_B_INPUT"] = uint8(y)
}
