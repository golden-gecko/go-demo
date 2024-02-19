#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml --project-name demo-stack \
    build --pull "$@"

docker-compose --file docker-compose.yml --project-name demo-stack \
    push "$@"
