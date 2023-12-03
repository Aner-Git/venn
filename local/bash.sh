#!/bin/sh

set -e

if [ -z "$1" ]; then
    echo "service name required - (e.g: meatmath)"
    return 1
fi

SERVICE_NAME=$1
shift
./compose.sh exec "$SERVICE_NAME" bash
