#!/bin/bash

if [ -z "$1" ]; then
    echo "rewrite template file required - (e.g: index.html.template)"
    exit 1
fi

if [ -z "$2" ]; then
    echo "template result file required - (e.g: index.html)"
    exit 1
fi

REWRITE_TEMPLATE_FILE=$1
REWRITE_RESULT_FILE=$2
OVERWRITE=$3

if [ "$OVERWRITE" == "-o" ] || [ ! -f "$REWRITE_RESULT_FILE" ]; then
    cp "$REWRITE_TEMPLATE_FILE" "$REWRITE_RESULT_FILE"
fi

for var in "${!BOOT_@}"; do
    sed -i "s|\"__${var}__\"|\"${!var}\"|g" "$REWRITE_RESULT_FILE"
done

for var in "${!REACT_APP_@}"; do
    sed -i "s|\"__${var}__\"|\"${!var}\"|g" "$REWRITE_RESULT_FILE"
done
