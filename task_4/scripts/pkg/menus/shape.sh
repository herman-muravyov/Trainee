#! /bin/bash

function shape_menu()
{
	clear
	echo -e "\n$(set_color "purple")----------------------------------------------------------------------------------------------$(set_color "*")"
	echo -e "\nShape Menu:"
	echo "$(set_color "purple")[1]$(set_color "*") - Circle"
	echo "$(set_color "purple")[2]$(set_color "*") - Rectangle"
	echo "$(set_color "purple")[0]$(set_color "*") <- Back"
}