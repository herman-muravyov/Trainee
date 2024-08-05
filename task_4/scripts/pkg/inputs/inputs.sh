#! /bin/bash


function menu_input()
{
	read -p "Enter your choice: " choice
}

function shape_circle_params() {
	read -p "Enter the radius value: " radius
}

function shape_rectangle_width() {
	read -p "Enter the width value: " width
}

function shape_rectangle_height() {
	read -p "Enter the height value: " height
}

function continue_using()
{
	echo "$(set_color "green")Press any button to continue...$(set_color "*")"
	read -r
}