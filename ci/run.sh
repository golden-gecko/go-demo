#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml --project-name demo-ci \
    up --build --detach --remove-orphans "$@"
