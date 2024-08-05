#! /bin/bash

source ./scripts/assets/set_color.sh
source ./scripts/lib/services/helpers/download.sh
source ./scripts/pkg/inputs/inputs.sh
source ./scripts/pkg/errors/errors.sh
source ./scripts/pkg/menus/services/yandex.sh

function yandex_service()
{
	while true; do
		yandex_menu
		menu_input

		case $choice in
			1)
                url="https://data-sets.storage.yandexcloud.net/sets/coffee-flavors.csv"
                download_file "$url" 
                ;;
			2)
                url="https://data-sets.storage.yandexcloud.net/sets/unemployment_data_europe_oecd.csv"
                download_file "$url" 
                ;;
            3)
                url="https://data-sets.storage.yandexcloud.net/sets/european_leaders.csv"
                download_file "$url" 
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