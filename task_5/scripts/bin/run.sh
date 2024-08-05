#! /bin/bash

source ./scripts/pkg/menus/main.sh
source ./scripts/pkg/inputs/inputs.sh
source ./scripts/lib/services/services.sh


function main()
{
	while True; do
		main_menu
		menu_input

		case $choice in
			1)
				services
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