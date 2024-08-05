package utils

import (
	"flag"
	"log"
	"strconv"
)

type Flags struct {
	Shape  string
	Radius float64
	Width  float64
	Height float64
}

func ParseFlags() Flags {
	var shape string
	var radius, width, height string

	flag.StringVar(&shape, "shape", "", "The shape to calculate the area for (circle or rectangle)")
	flag.StringVar(&radius, "radius", "0", "The radius of the circle")
	flag.StringVar(&width, "width", "0", "The width of the rectangle")
	flag.StringVar(&height, "height", "0", "The height of the rectangle")
	flag.Parse()

	radiusValue, err := strconv.ParseFloat(radius, 64)
	if err != nil {
		log.Fatalf("Invalid value for -radius: %s, - is should be type INT", radius)
	}

	widthValue, err := strconv.ParseFloat(width, 64)
	if err != nil {
		log.Fatalf("Invalid value for -width: %s, - is should be type INT", width)
	}

	heightValue, err := strconv.ParseFloat(height, 64)
	if err != nil {
		log.Fatalf("Invalid value for -height: %s, - is should be type INT", height)
	}

	return Flags{
		Shape:  shape,
		Radius: radiusValue,
		Width:  widthValue,
		Height: heightValue,
	}
}
