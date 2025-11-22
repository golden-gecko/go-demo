#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml --file docker-compose-elk.yml \
    up --build --detach --remove-orphans "$@"
