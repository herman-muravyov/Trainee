#! /bin/bash

function output_validator()
{
    local output=$1

    if [[ -z "$output" ]]; then
		echo "$(set_color "red")File name cannot be empty.$(set_color "*")" >&2
		return 1
  	fi
}