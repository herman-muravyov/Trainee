#! /bin/bash

source ./scripts/pkg/inputs/inputs.sh
source ./scripts/lib/services/helpers/download.sh
source ./scripts/pkg/validators/url_validator.sh

function custom_url() {
    url_input
    if ! is_url "$url"; then
        echo "Invalid URL provided. Please try again."
        continue_using
    else
        download_file "$url"
    fi
}