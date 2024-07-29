#! /bin/bash

source ./scripts/assets/set_color.sh
source ./scripts/pkg/inputs/inputs.sh

function is_url()
{
	local url="$1"

	if [[ $url =~ ^(https?|ftp):// ]]; then
		return 0
	else
		echo "$(set_color "red")Error.$(set_color "*") Invalid URL." >&2
		return 1
	fi
}