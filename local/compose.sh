#!/bin/sh

set -e

DOCKER_BUILDKIT=1 docker-compose \
--project-name=venn \
--project-directory=.. \
--file=./docker-compose.yml \
"$@"
