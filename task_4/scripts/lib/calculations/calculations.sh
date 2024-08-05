#! /bin/bash

source ./scripts/pkg/menus/shape.sh
source ./scripts/pkg/inputs/inputs.sh
source ./scripts/lib/calculations/circle.sh
source ./scripts/assets/set_color.sh
source ./scripts/lib/calculations/rectangle.sh


function calculations()
{
	while true; do
		shape_menu
		menu_input

 		echo "Choice selected: $choice"

		case $choice in
			1)
				circle_calc
				continue_using
				;;
			2)
				rectangle_calc
				continue_using
				;;
			0)
				break
				;;
			*)
				invalid_menu_input	
				continue_using
				;;
		esac
	done
}