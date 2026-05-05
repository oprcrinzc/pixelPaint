// Package state for state managing
package state

import (
	"pixelpaint/data"
)

// --------------------------- State struct ----------------------------------------------

type State struct {
	main     func(s *State)
	load     func(s *State)
	unload   func(s *State)
	isLoaded bool
	Name     string
	Storage  map[string]any
}

func (s *State) New(name string) *State {
	s.isLoaded = false
	s.Name = name
	s.Storage = make(map[string]any)
	return s
}

func (s *State) SetIsLoaded(b bool) *State {
	s.isLoaded = b
	return s
}

func (s *State) SetLoadFunc(f func(s *State)) *State {
	s.load = f
	return s
}

func (s *State) SetUnloadFunc(f func(s *State)) *State {
	s.unload = f
	return s
}

func (s *State) SetMainFunc(f func(s *State)) *State {
	s.main = f
	return s
}

func (s *State) Main() {
	s.main(s)
}

func (s *State) Load() {
	s.load(s)
}

func (s *State) UnLoad() {
	s.unload(s)
}

// -------------------------------------------------------------------------

var StateList map[string]*State = make(map[string]*State)

var (
	StateCreateSheet *State = new(State)
	StatePaint       *State = new(State)
)

// -------------------------------------------------------------------------

func Load() {
	StateCreateSheet.New("CreateSheet").
		SetLoadFunc(StateCreateSheetLoad).
		SetMainFunc(StateCreateSheetMain).
		SetUnloadFunc(StateCreateSheetUnLoad)
	StateList[StateCreateSheet.Name] = StateCreateSheet

	StatePaint.New("Paint").
		SetLoadFunc(StatePaintLoad).
		SetMainFunc(StatePaintMain).
		SetUnloadFunc(StatePaintUnload)
	StateList[StatePaint.Name] = StatePaint
}

func Run() {
	s, ok := StateList[data.CurrentState]
	if ok {
		s.Main()
	}
}
