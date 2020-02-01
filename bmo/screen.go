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
func (s *Screen) Position(x, y int32) Point {
	return s.position(x, y)
}

func (s *Screen) TouchPosition(x, y float32) Point {
	return s.Position(
		int32(x * float32(s.Rect.W)),
		int32(y * float32(s.Rect.H)),
	)
}

func (p Point) IsInside(rect *sdl.Rect) bool {
	return rect.X <= p.X && p.X <= rect.X + rect.W - 1 &&
		rect.Y <= p.Y && p.Y <= rect.Y + rect.H - 1
}
