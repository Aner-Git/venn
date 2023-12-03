#!/bin/sh

set -e

# Check if docker COMPOSE_PROFILES is passed
if [ -z "${COMPOSE_PROFILES}" ]; then
echo "Running minimal services, set COMPOSE_PROFILES=all to run all services"
else
export COMPOSE_PROFILES
fi

./compose.sh \
up \
-d \
--remove-orphans \
"$@"

./compose-logs.sh
