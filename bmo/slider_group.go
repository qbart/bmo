package bmo

import "github.com/veandco/go-sdl2/sdl"

// Component struct.
type SliderGroup struct {
	c *Component
}

func NewSliderGroup(rect sdl.Rect, color RGBColor) *SliderGroup {
	sg := &SliderGroup{
		c: NewComponent(rect, color),
	}
	sg.OnMouseMoved(sg.mouseMoveCallback)
	return sg
}

func (sg *SliderGroup) Draw(renderer *sdl.Renderer) {
	if sg.c.visible {
		renderer.SetDrawColor(
			sg.c.Color.R,
			sg.c.Color.G,
			sg.c.Color.B,
			sg.c.Color.A,
		)
		renderer.FillRect(&sg.c.Rect)
	}
}

func (sg *SliderGroup) Contains(point Point) bool {
	return sg.c.Contains(point)
}

func (sg *SliderGroup) Show(visible bool) {
	sg.c.Show(visible)
}

func (sg *SliderGroup) OnMousePressed(callback OnMouseEventCallback) {
	sg.c.OnMousePressed(callback)
}

func (sg *SliderGroup) OnMouseMoved(callback OnMouseEventCallback) {
	sg.c.OnMouseMoved(callback)
}

func (sg *SliderGroup) OnMouseReleased(callback OnMouseEventCallback) {
	sg.c.OnMouseReleased(callback)
}

func (sg *SliderGroup) TriggerOnMousePressed(p Point) {
	sg.c.TriggerOnMousePressed(p)
}

func (sg *SliderGroup) TriggerOnMouseMoved(p Point) {
	sg.c.TriggerOnMouseMoved(p)
}

func (sg *SliderGroup) TriggerOnMouseReleased(p Point) {
	sg.c.TriggerOnMouseReleased(p)
}

func (sg *SliderGroup) mouseMoveCallback(e MouseEvent) {
	//TODO
}
