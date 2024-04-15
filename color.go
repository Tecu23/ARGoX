package main

// Color represent either WHITE or BLACK
type Color int

// Opp should return the opposite color
func (c Color) Opp() Color {
	return c ^ 0x1
}

// String returns the string representation of the color
func (c Color) String() string {
	if c == WHITE {
		return "W"
	}
	return "B"
}
