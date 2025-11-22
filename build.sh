#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml \
    pull "$@"

docker-compose --file docker-compose.yml \
    build "$@"
