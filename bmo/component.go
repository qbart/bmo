package bmo

import "github.com/veandco/go-sdl2/sdl"

type MouseEvent struct {
	P Point
}

type OnMouseEventCallback func(event MouseEvent)

type IComponent interface {
	Draw(renderer *sdl.Renderer)
	Contains(point Point) bool
	Show(visible bool)

	// mouse support
	OnMousePressed(callback OnMouseEventCallback)
	OnMouseMoved(callback OnMouseEventCallback)
	OnMouseReleased(callback OnMouseEventCallback)
	TriggerOnMousePressed(p Point)
	TriggerOnMouseMoved(p Point)
	TriggerOnMouseReleased(p Point)
}

// Component struct.
type Component struct {
	Rect sdl.Rect
	Color RGBColor
	visible bool
	onMousePressedListeners []OnMouseEventCallback
	onMouseMovedListeners []OnMouseEventCallback
	onMouseReleasedListeners []OnMouseEventCallback
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
		onMousePressedListeners: make([]OnMouseEventCallback, 0),
		onMouseMovedListeners: make([]OnMouseEventCallback, 0),
		onMouseReleasedListeners: make([]OnMouseEventCallback, 0),
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

func (c *Component) OnMousePressed(callback OnMouseEventCallback) {
	c.onMousePressedListeners = append(c.onMousePressedListeners, callback)
}

func (c *Component) OnMouseMoved(callback OnMouseEventCallback) {
	c.onMouseMovedListeners = append(c.onMouseMovedListeners, callback)
}

func (c *Component) OnMouseReleased(callback OnMouseEventCallback) {
	c.onMouseReleasedListeners = append(c.onMouseReleasedListeners, callback)
}

func (c *Component) TriggerOnMousePressed(p Point) {
	c.handleMouseEvent(p, &c.onMousePressedListeners)
}

func (c *Component) TriggerOnMouseMoved(p Point) {
	c.handleMouseEvent(p, &c.onMouseMovedListeners)
}

func (c *Component) TriggerOnMouseReleased(p Point) {
	c.handleMouseEvent(p, &c.onMouseReleasedListeners)
}

func (c *Component) handleMouseEvent(p Point, listeners *[]OnMouseEventCallback) {
	if c.visible && c.Contains(p) {
		relativeP := Point{
			p.X - c.Rect.X,
			p.Y - c.Rect.Y,
		}
		for _, listener := range *listeners {
			listener(MouseEvent{
				P: relativeP,
			})
		}
	}
}
