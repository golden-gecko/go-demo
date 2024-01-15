#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml \
    ps --all "$@"
