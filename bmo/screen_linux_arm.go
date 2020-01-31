package bmo

func (s *Screen) position(event *sdl.MouseButtonEvent) (int32, int32) {
	return event.X, s.H - event.Y - 1
}
