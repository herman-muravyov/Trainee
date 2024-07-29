#! /bin/bash

source ./scripts/assets/set_color.sh
source ./scripts/assets/ascii.sh

function main_menu()
{
	clear
	display_art
	echo -e "\n$(set_color "purple")-----------------------------------------------------------------------------\
-----------------$(set_color "*")"
	echo -e "\nMenu:"
	echo "$(set_color "purple")[1]$(set_color "*") - Calculate a shape"
	echo "$(set_color "purple")[0]$(set_color "*") - Exit"
}