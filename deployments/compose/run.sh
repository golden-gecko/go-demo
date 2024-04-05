#!/bin/bash -ex

cd "$(dirname "$0")"

docker volume create demo-roach-certs

docker-compose --file docker-compose.yml --project-name demo-src \
    up --build --detach --remove-orphans "$@"
