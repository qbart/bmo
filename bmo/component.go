package bmo

import "github.com/veandco/go-sdl2/sdl"

type MouseEvent struct {
	P Point
}

type OnMousePressedCallback func(event MouseEvent)

type IComponent interface {
	Draw(renderer *sdl.Renderer)
	Contains(point Point) bool
	Show(visible bool)

	// mouse support
	OnMousePressed(callback OnMousePressedCallback)
	TriggerOnMousePressed(p Point)
}

// Component struct.
type Component struct {
	Rect sdl.Rect
	Color RGBColor
	visible bool
	listeners []OnMousePressedCallback
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

func NewComponent(rect sdl.Rect, color RGBColor) *Component {
	return &Component{
		Rect: rect,
		Color: color,
		visible: false,
		listeners: make([]OnMousePressedCallback, 0),
	}
}

func (c *Component) Draw(renderer *sdl.Renderer) {
	if c.visible {
		renderer.SetDrawColor(c.Color.R, c.Color.G, c.Color.B, c.Color.A)
		renderer.FillRect(&c.Rect)
	}
}

func (c *Component) Contains(point Point) bool {
	return c.visible &&
		c.Rect.X <= point.X && point.X <= c.Rect.X + c.Rect.W - 1 &&
		c.Rect.Y <= point.Y && point.Y <= c.Rect.Y + c.Rect.H - 1
}

func (c *Component) Show(visible bool) {
	c.visible = visible
}

func (c *Component) OnMousePressed(callback OnMousePressedCallback) {
	c.listeners = append(c.listeners, callback)
}

func (c *Component) TriggerOnMousePressed(p Point) {
	relativeP := Point{
		p.X - c.Rect.X,
		p.Y - c.Rect.Y,
	}
	for _, listener := range c.listeners {
		listener(MouseEvent{
			P: relativeP,
		})
	}
}
