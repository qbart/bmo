package bmo

import "github.com/veandco/go-sdl2/sdl"

type IComponent interface {
	Draw(renderer *sdl.Renderer)
	Contains(point Point) bool
}

// Component struct.
type Component struct {
	Rect sdl.Rect
	Color RGBColor
}

type RGBColor struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func RGB(r, g, b uint8) RGBColor {
	return RGBColor{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func (c *Component) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(c.Color.R, c.Color.G, c.Color.B, c.Color.A)
	renderer.FillRect(&c.Rect)
}

func (c *Component) Contains(point Point) bool {
	return c.Rect.X <= point.X && point.X <= c.Rect.X + c.Rect.W - 1 &&
		c.Rect.Y <= point.Y && point.Y <= c.Rect.Y + c.Rect.H - 1
}
