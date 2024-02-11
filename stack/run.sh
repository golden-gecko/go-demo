#!/bin/bash -ex

cd "$(dirname "$0")"

docker-compose --file docker-compose.yml --project-name demo-stack \
    up --build --detach --remove-orphans "$@"
