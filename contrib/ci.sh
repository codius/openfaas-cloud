#!/bin/sh
set -e
# Run this command from one-level higher in the folder path, not this folder.

CLI="faas-cli"
if ! [ -x "$(command -v faas-cli)" ]; then
    HERE=`pwd`
    cd /tmp/
    curl -sL https://github.com/openfaas/faas-cli/releases/download/0.9.3/faas-cli > faas-cli
    chmod +x ./faas-cli
    CLI="/tmp/faas-cli"

    echo "Returning to $HERE"
    cd $HERE
fi

echo "Working folder: `pwd`"

$CLI build --parallel=4
HERE=`pwd`
cd dashboard
make build-dist
$CLI build
cd $HERE