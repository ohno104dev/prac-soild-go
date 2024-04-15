package bad

import "math"

type Shape struct {
	Type   string
	Width  float64
	Height float64
	Radius float64
}

func CalculateArea(s Shape) float64 {
	if s.Type == "rectangle" {
		return s.Width * s.Height
	} else if s.Type == "circle" {
		return math.Pi * s.Radius * s.Radius
	}

	// adding support for a new shape would require modifying the CalculateArea function

	return 0
}
