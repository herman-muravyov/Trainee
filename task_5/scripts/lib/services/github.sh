#! /bin/bash

source ./scripts/assets/set_color.sh
source ./scripts/lib/services/helpers/download.sh
source ./scripts/pkg/menus/services/github.sh
source ./scripts/pkg/inputs/inputs.sh
source ./scripts/pkg/errors/errors.sh

function github_service()
{
	while true; do
		github_menu
		menu_input

		case $choice in
			1)
                url="https://raw.githubusercontent.com/plotly/datasets/master/stockdata.csv"
				download_file "$url" 
				;;
			2)
                url="https://raw.githubusercontent.com/plotly/datasets/master/stockdata2.csv"
				download_file "$url" 
				;;
            3)
                url="https://raw.githubusercontent.com/plotly/datasets/master/tesla-stock-price.csv"
				download_file "$url" 
				;;
            4)
                url="https://raw.githubusercontent.com/plotly/datasets/master/2014_us_cities.csv"
				download_file "$url" 
				;;
            5)
                url="https://raw.githubusercontent.com/plotly/datasets/master/2014_usa_states.csv"
				download_file "$url" 
				;;
            6)
                url="https://raw.githubusercontent.com/plotly/datasets/master/dash-stock-ticker-demo.csv"
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