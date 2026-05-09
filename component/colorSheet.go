package component

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ColorSheet struct {
	Width         int32
	Height        int32
	colorMap      map[int32]map[int32]color.RGBA
	renderTexture rl.RenderTexture2D
	onRender      func(c *ColorSheet)
}

func (c *ColorSheet) New(w, h int32) *ColorSheet {
	c.Height = h
	c.Width = w
	c.renderTexture = rl.LoadRenderTexture(w, h)
	c.colorMap = make(map[int32]map[int32]color.RGBA, h)
	for i := range h {
		c.colorMap[i] = make(map[int32]color.RGBA, w)
	}
	return c
}

func (c *ColorSheet) Set(x, y int32, col color.RGBA) *ColorSheet {
	c.colorMap[y][x] = col
	rl.DrawPixel(x, y, col)
	return c
}

func (c *ColorSheet) Bind(f func(c *ColorSheet)) *ColorSheet {
	c.onRender = f
	return c
}

func (c *ColorSheet) Render() rl.Texture2D {
	rl.BeginTextureMode(c.renderTexture)
	c.onRender(c)
	rl.EndTextureMode()
	return c.renderTexture.Texture
}
