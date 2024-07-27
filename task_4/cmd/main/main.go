package main

import (
	"fmt"
	"log"
	"task/calculations"
	"task/interfaces"
	"task/utils"
)

func main() {
	var s interfaces.Shape
	flags := utils.ParseFlags()

	if flags.Shape == "" {
		log.Fatal("Error: -shape flag is required")
	}

	switch flags.Shape {
	case "circle":
		if flags.Radius <= 0 {
			log.Fatal("Error: -radius must be greater than 0 for circle")
		}
		s = &calculations.Circle{Radius: flags.Radius}
	case "rectangle":
		if flags.Width <= 0 {
			log.Fatal("Error: -width must be greater than 0 for rectangle")
		}
		if flags.Height <= 0 {
			log.Fatal("Error: -height must be greater than 0 for rectangle")
		}
		s = &calculations.Rectangle{Width: flags.Width, Height: flags.Height}
	default:
		log.Fatal("Error: invalid shape. Must be 'circle' or 'rectangle'")
	}

	fmt.Printf("The area of the %s is: %.2f\n", flags.Shape, s.Area())
}
