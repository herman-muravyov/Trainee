#! /bin/bash

source ./scripts/pkg/inputs/inputs.sh
source ./scripts/pkg/helpers/success_output.sh
source ./scripts/pkg/validators/output_validator.sh
source ./scripts/pkg/validators/extension_validator.sh

function download_file() {
    local url=$1

    output_input
    if ! output_validator "$output"; then
        continue_using
    fi

    file_extension_input
    if ! extension_validator "$ext"; then
        continue_using
    fi
    
    go run cmd/main/main.go -url "$url" -output "$output""$ext"
    success_output
    continue_using
}