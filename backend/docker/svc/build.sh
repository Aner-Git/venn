#!/bin/sh

set -e

if [ -z "$1" ] || [ -z "$2" ] ; then
    echo "a branch name is required - (e.g: main)"
    echo "a tag/version name is required - (e.g: a,b,c ) would be appended with date"
    exit 1
fi

APP_NAME=meatmath
BRANCH=$1
TAG=${2}`date +%d%m%g`

echo "using tag: ${TAG}" 

TARGET=production
TAG_LATEST="$BRANCH-latest"

(
cd ../.. || exit 1
DOCKER_BUILDKIT=1 docker build \
--rm=true \
--target "$TARGET" \
-t "$APP_NAME":"$TAG" \
-t "$APP_NAME":"$TAG_LATEST" \
-f docker/svc/Dockerfile . \
--build-arg service_name="$APP_NAME" \
--build-arg tag="$TAG" \
--build-arg git_commit="<to-do>" \
)
