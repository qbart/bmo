package bmo

import "github.com/veandco/go-sdl2/sdl"

type Point struct {
	X int32
	Y int32
}

type Screen struct {
	Rect sdl.Rect
}

// Position calculates screen position.
func (s *Screen) Position(event *sdl.MouseButtonEvent) Point {
	return s.position(event)
}
