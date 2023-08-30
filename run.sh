#!/usr/bin/env bash

# Download EOWL
EOWL_URL="https://diginoodles.com/projects/eowl/EOWL-v1.1.2.zip"
if [ ! -d "EOWL-v1.1.2" ]; then
    # If unzip is not installed - show error
    if ! [ -x "$(command -v unzip)" ]; then
        echo "Error: 'unzip' command is required to run this script" >&2
        exit 1
    fi
    curl -L -o EOWL-v1.1.2.zip $EOWL_URL
    unzip EOWL-v1.1.2.zip "EOWL-v1.1.2/LF Delimited Format/*Words.txt" -d . > /dev/null
    rm EOWL-v1.1.2.zip
    # Combine all words into one file
    cat EOWL-v1.1.2/"LF Delimited Format"/*Words.txt > EOWL-v1.1.2/words.txt
    # Remove all other files
    rm -rf EOWL-v1.1.2/"LF Delimited Format"
fi

go get .
go run .
