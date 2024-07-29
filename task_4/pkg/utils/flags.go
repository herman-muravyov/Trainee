package utils

import "flag"

type Flags struct {
	Shape  string
	Width  float64
	Height float64
	Radius float64
}

func ParseFlags() Flags {
	shape := flag.String("shape", "", "The shape to calculate the area for (circle or rectangle)")
	width := flag.Float64("width", 0, "The width of the rectangle")
	height := flag.Float64("height", 0, "The height of the rectangle")
	radius := flag.Float64("radius", 0, "The radius of the circle")

	flag.Parse()

	return Flags{
		Shape:  *shape,
		Width:  *width,
		Height: *height,
		Radius: *radius,
	}
}
