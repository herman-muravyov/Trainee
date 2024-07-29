#! /bin/bash

source ./scripts/pkg/menus/main.sh
source ./scripts/pkg/inputs/inputs.sh
source ./scripts/lib/calculations/calculations.sh


function main()
{
	while True; do
		main_menu
		menu_input

		case $choice in
			1)
				calculations
				;;
			0)
				echo "$(set_color "purple")Goodbye!$(set_color "*")"
				return 1
				;;
			*)
				invalid_menu_input	
				continue_using
				;;
		esac
	done
}

main