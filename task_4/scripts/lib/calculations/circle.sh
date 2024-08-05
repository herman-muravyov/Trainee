#! /bin/bash

source ./scripts/pkg/inputs/inputs.sh


function circle_calc()
{
	local radius
	
	shape_circle_params

	go run cmd/main/main.go -shape circle -radius "$radius"
}