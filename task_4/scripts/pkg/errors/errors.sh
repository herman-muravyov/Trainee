#! /bin/bash

function invalid_menu_input()
{
	echo "$(set_color "red")Envalid choice.$(set_color "*") Try again." >&2
}