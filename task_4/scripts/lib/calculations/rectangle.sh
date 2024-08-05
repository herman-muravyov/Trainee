#! /bin/bash

source ./scripts/pkg/inputs/inputs.sh

function rectangle_calc()
{
	local width height
	
	shape_rectangle_width
	shape_rectangle_height

	go run cmd/main/main.go -shape rectangle -width "$width" -height "$height"
}