#!/bin/bash -ex

cd "$(dirname "$0")"

export COMPOSE_DOCKER_CLI_BUILD=0
export DOCKER_BUILDKIT=0

docker-compose --file docker-compose.yml --project-name demo-stack \
    build --pull "$@"

docker-compose --file docker-compose.yml --project-name demo-stack \
    push "$@"
