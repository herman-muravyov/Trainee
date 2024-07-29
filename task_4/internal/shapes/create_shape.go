package shapes

import (
	"log"
	"task/internal/calculations"
	"task/internal/interfaces"
	"task/pkg/utils"
)

func CreateShape(flags utils.Flags) interfaces.Shape {
	switch flags.Shape {
	case "circle":
		return &calculations.Circle{Radius: flags.Radius}
	case "rectangle":
		return &calculations.Rectangle{Width: flags.Width, Height: flags.Height}
	default:
		log.Fatal("Error: invalid shape. Must be 'circle' or 'rectangle'")
		return nil
	}
}
