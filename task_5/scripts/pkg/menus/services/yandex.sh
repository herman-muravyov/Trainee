#! /bin/bash

source ./scripts/assets/set_color.sh

function yandex_menu()
{
	clear
	display_art
	echo -e "\n$(set_color "purple")----------------------------------------------------------------------------------------------$(set_color "*")"
	echo -e "\nYandex Cloud Menu:"
	echo "$(set_color "purple")[1]$(set_color "*") - Coffee Flavors"
	echo "$(set_color "purple")[2]$(set_color "*") - Unemployment Rate Europe"
    echo "$(set_color "purple")[3]$(set_color "*") - European Leaders"
	echo "$(set_color "purple")[0]$(set_color "*") <- Back"
}