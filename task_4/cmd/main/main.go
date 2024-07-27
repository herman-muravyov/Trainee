package main

import (
	"fmt"
	"task/shapes"
	"task/utils"
	"task/validators"
)

func main() {
	flags := utils.ParseFlags()
	validators.ValidateFlags(flags)
	shape := shapes.CreateShape(flags)

	fmt.Printf("The area of the %s is: %.2f\n", flags.Shape, shape.Area())
}
