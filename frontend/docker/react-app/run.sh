#!/bin/bash

set -ex

echo "Rewriting index.html"
./rewrite-index.sh index.html.template index.html

# execute caddy in the background, and start a terminal. 
exec su -c "exec /webapp/caddy run --config /webapp/Caddyfile" -s /bin/sh nobody
