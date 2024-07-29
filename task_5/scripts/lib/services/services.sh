#! /bin/bash

source ./scripts/assets/set_color.sh
source ./scripts/lib/services/github.sh
source ./scripts/lib/services/yandex.sh
source ./scripts/lib/services/custom.sh
source ./scripts/pkg/inputs/inputs.sh
source ./scripts/pkg/menus/services/services.sh
source ./scripts/pkg/errors/errors.sh

function services()
{
	while true; do
		services_menu
		menu_input

		case $choice in
			1)
                github_service
				;;
			2)
                yandex_service
				;;
			3)
                custom_url
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