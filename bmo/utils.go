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
