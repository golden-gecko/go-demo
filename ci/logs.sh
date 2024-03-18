#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml  --project-name demo-ci \
    logs --follow --tail=20 "$@"
