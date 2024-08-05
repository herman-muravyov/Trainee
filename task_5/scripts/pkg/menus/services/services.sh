#! /bin/bash

source ./scripts/assets/set_color.sh

function services_menu()
{
	clear
	display_art
	echo -e "\n$(set_color "purple")----------------------------------------------------------------------------------------------$(set_color "*")"
	echo -e "\nServices Menu:"
	echo "$(set_color "purple")[1]$(set_color "*") - GitHub"
	echo "$(set_color "purple")[2]$(set_color "*") - YandexCloud"
	echo "$(set_color "purple")[3]$(set_color "*") - Custom URL"
	echo "$(set_color "purple")[0]$(set_color "*") <- Back"
}