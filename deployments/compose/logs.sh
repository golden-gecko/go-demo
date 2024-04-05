#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml  --project-name demo-src \
    logs --follow --tail=20 "$@"
