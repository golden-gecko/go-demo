#!/bin/bash -ex

cd "$(dirname "$0")"

export COMPOSE_DOCKER_CLI_BUILD=0
export DOCKER_BUILDKIT=0

docker-compose --file docker-compose.yml --project-name demo-stack \
    up --build --detach --remove-orphans "$@"
