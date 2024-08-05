#! /bin/bash

source ./scripts/assets/set_color.sh

function extension_validator() {
    local ext="$1"
    if [[ -z "$ext" ]]; then
        echo "$(set_color "red")Extension cannot be empty. Please enter a valid file extension.$(set_color "*")"
        file_extension_input
    elif [[ "$ext" =~ ^\.[a-z]+$ ]]; then
        echo "Valid extension: $ext"
    else
        echo "$(set_color "red")Invalid extension. Please enter a valid file extension starting with a dot and followed by lowercase alphabetic characters, with no numbers or uppercase letters.$(set_color "*")"
        file_extension_input
    fi
}