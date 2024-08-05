package validators

import (
	"log"
	"task/pkg/utils"
)

func ValidateFlags(flags utils.Flags) {
	switch flags.Shape {
	case "circle":
		if flags.Radius <= 0 || (flags.Width != 0 || flags.Height != 0) {
			log.Fatal("Error: for circle, -radius must be greater than 0 and -width/-height should not be set")
		}
	case "rectangle":
		if flags.Width <= 0 || flags.Height <= 0 || flags.Radius != 0 {
			log.Fatal("Error: for rectangle, -width and -height must be greater than 0 and -radius should not be set")
		}
	default:
		log.Fatal("Error: invalid shape. Must be 'circle' or 'rectangle'")
	}
}
