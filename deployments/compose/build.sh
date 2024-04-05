#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml --project-name demo-src \
    build --pull "$@"

docker-compose --file docker-compose.yml --project-name demo-src \
    push "$@"
