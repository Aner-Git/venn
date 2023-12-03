#!/bin/sh

set -e

if [ -z "$1" ]; then
    echo "service name required - (e.g: meathmath)"
    return 1
fi

SERVICE_NAME=$1
shift
./compose.sh exec "$SERVICE_NAME"-mysql /usr/bin/mysql -u "$SERVICE_NAME" -p"$SERVICE_NAME" "$SERVICE_NAME" "$@"
