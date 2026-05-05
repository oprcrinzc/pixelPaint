package state

import (
	"strconv"

	"pixelpaint/component"
	"pixelpaint/data"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func StateCreateSheetLoad(s *State) {
	inputWidthPx := component.InputText{}
	inputWidthPx.New(300, rl.NewVector2(300, 50))

	inputHeightPx := component.InputText{}
	inputHeightPx.New(300, rl.NewVector2(300, 120))

	s.Storage["inputWidth"] = &inputWidthPx
	s.Storage["inputHeight"] = &inputHeightPx

	createBtn := component.Button{}
	createBtn.New("Create", rl.NewVector2(300, 190)).
		Bind(func() {
			var width string
			var height string
			if i, ok := s.Storage["inputWidth"]; ok {
				j := i.(*component.InputText)
				width = j.GetValue()
			}
			if i, ok := s.Storage["inputHeight"]; ok {
				j := i.(*component.InputText)
				height = j.GetValue()
			}
			if w, err := strconv.Atoi(width); err == nil {
				data.SheetWidth = w
			}
			if h, err := strconv.Atoi(height); err == nil {
				data.SheetHeight = h
			}
		})
	s.Storage["createBtn"] = &createBtn

	s.SetIsLoaded(true)
}

func StateCreateSheetUnLoad(s *State) {}

// -------------------------------------------------------------------------

func StateCreateSheetMain(s *State) {
	if !s.isLoaded {
		s.Load()
	}
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Width", 100, 50, 50, rl.Black)
	rl.DrawText("Height", 100, 120, 50, rl.Black)
	if iw, ok := s.Storage["inputWidth"]; ok {
		iw.(*component.InputText).Render()
	}
	if ih, ok := s.Storage["inputHeight"]; ok {
		ih.(*component.InputText).Render()
	}

	if btn, ok := s.Storage["createBtn"]; ok {
		btn.(*component.Button).Render()
	}
}
