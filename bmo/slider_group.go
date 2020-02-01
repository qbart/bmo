package bmo

import "github.com/veandco/go-sdl2/sdl"

// Component struct.
type SliderGroup struct {
	c      *Component
	values [4]float32
	active int
}

func NewSliderGroup(rect sdl.Rect, color RGBColor) *SliderGroup {
	sg := &SliderGroup{
		c: NewComponent(rect, color),
		values: [4]float32{1, 0, 0.5, 0.75},
		active: -1,
	}
	sg.OnMousePressed(sg.mousePressCallback)
	sg.OnMouseMoved(sg.mouseMoveCallback)
	sg.OnMouseReleased(sg.mouseReleaseCallback)
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
		renderer.SetDrawColor(5, 59, 43, 255)
		for i := 0; i < len(sg.values); i++ {
			rect := sg.colRect(i)
			renderer.FillRect(&rect)
		}
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

func (sg *SliderGroup) mousePressCallback(e MouseEvent) {
	sg.active = -1
	for i := 0; i < len(sg.values); i++ {
		rect := sg.colMaxRect(i)
		if e.GP.IsInside(&rect) {
			sg.active = i
			break
		}
	}
}

func (sg *SliderGroup) mouseMoveCallback(e MouseEvent) {
	if sg.active != -1 {
		rect := sg.colMaxRect(sg.active)
		x := e.GP.Y - rect.Y
		p := float32(x) / float32(rect.H)
		sg.values[sg.active] = Clampf(1 - p, 0, 1)
	}
}

func (sg *SliderGroup) mouseReleaseCallback(e MouseEvent) {
	sg.active = -1
}

func (sg *SliderGroup) colOffset() Point {
	return Point{sg.c.Rect.X+10, sg.c.Rect.Y+11}
}

func (sg *SliderGroup) colMaxHeight() int32 {
	return int32(180)
}

func (sg *SliderGroup) colRect(i int) sdl.Rect {
	offset := sg.colOffset()
	h := int32(float32(sg.colMaxHeight()) * sg.values[i])
	return sdl.Rect{
		offset.X + 58 * int32(i),
		offset.Y + sg.colMaxHeight() - h,
		45,
		h,
	}
}

func (sg *SliderGroup) colMaxRect(i int) sdl.Rect {
	offset := sg.colOffset()
	return sdl.Rect{
		offset.X + 58 * int32(i),
		offset.Y,
		45,
		sg.colMaxHeight(),
	}
}
