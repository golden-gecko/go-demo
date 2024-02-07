#!/bin/bash -ex

cd "$(dirname "$0")"

export COMPOSE_DOCKER_CLI_BUILD=0
export DOCKER_BUILDKIT=0

hostname=4450ed6e2380 # $(docker exec -it demo-stack-rabbit-1-1 hostname)

docker exec -it demo-stack-rabbit-2-1 rabbitmqctl stop_app
docker exec -it demo-stack-rabbit-2-1 rabbitmqctl reset
docker exec -it demo-stack-rabbit-2-1 rabbitmqctl join_cluster rabbit@$hostname
docker exec -it demo-stack-rabbit-2-1 rabbitmqctl start_app

docker exec -it demo-stack-rabbit-3-1 rabbitmqctl stop_app
docker exec -it demo-stack-rabbit-3-1 rabbitmqctl reset
docker exec -it demo-stack-rabbit-3-1 rabbitmqctl join_cluster rabbit@$hostname
docker exec -it demo-stack-rabbit-3-1 rabbitmqctl start_app
