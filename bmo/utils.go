package bmo

// Clamp clamps "value" between a, b range.
func Clamp(value, a, b int) int {
	if value < a {
		value = a
	} else if value > b {
		value = b
	}

	return value
}

// Clamp clamps "value" between a, b range.
func Clampf(value, a, b float32) float32 {
	if value < a {
		value = a
	} else if value > b {
		value = b
	}

	return value
}
