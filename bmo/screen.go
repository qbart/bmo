package bmo

import "github.com/veandco/go-sdl2/sdl"

type Screen struct {
	W int32
	H int32
}

// Position calculates screen position.
func (s *Screen) Position(event *sdl.MouseButtonEvent) (int32, int32) {
	return s.position(event)
}
