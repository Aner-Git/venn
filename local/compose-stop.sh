#!/bin/sh

set -e

if [ -z "$1" ]; then
    echo "service name required - (e.g: api)"
    return 1
fi

./compose.sh stop "$1"
./compose.sh rm -f "$1"