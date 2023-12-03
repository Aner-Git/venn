#!/bin/sh

set -e

if [ -z "$1" ]; then
    echo "service name required - (e.g: api)"
    return 1
fi

./compose.sh build "$1"
./compose-stop.sh "$1"
./compose.sh up --no-start "$1"
./compose.sh restart "$1"

if [ ! "$2" = "--no-logs" ]; then
    ./compose-logs.sh "$1"
fi