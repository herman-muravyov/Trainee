package main

import (
	"fmt"
	"task/internal/shapes"
	"task/pkg/utils"
	"task/pkg/validators"
)

func main() {
	flags := utils.ParseFlags()
	validators.ValidateFlags(flags)
	shape := shapes.CreateShape(flags)

	fmt.Printf("The area of the %s is: %.2f\n", flags.Shape, shape.Area())
}
