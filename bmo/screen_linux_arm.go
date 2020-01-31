package bmo

import "github.com/veandco/go-sdl2/sdl"


func (s *Screen) position(event *sdl.MouseButtonEvent) Point {
	return Point{event.X, s.Rect.H - event.Y - 1}
}
