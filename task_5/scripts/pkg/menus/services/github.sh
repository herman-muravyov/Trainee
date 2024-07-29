#! /bin/bash

source ./scripts/assets/set_color.sh

function github_menu()
{
	clear
	display_art
	echo -e "\n$(set_color "purple")----------------------------------------------------------------------------------------------$(set_color "*")"
	echo -e "\nGitHub Menu:"
    echo "$(set_color "purple")[1]$(set_color "*") - Stock Data 1"
    echo "$(set_color "purple")[2]$(set_color "*") - Stock Data 2"
    echo "$(set_color "purple")[3]$(set_color "*") - Tesla Stock"
    echo "$(set_color "purple")[4]$(set_color "*") - 2014 US Cities"
	echo "$(set_color "purple")[5]$(set_color "*") - 2014 USA States"
    echo "$(set_color "purple")[6]$(set_color "*") - Dash Stock Ticker Demo"
	echo "$(set_color "purple")[0]$(set_color "*") <- Back"
}
