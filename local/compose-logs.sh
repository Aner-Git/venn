#!/bin/sh

set -e

./compose.sh logs --tail=1000 -f "$@"