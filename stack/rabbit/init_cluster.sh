#!/bin/bash -ex

cd "$(dirname "$0")"

name=$(docker exec -it demo-stack-rabbit-1-1 hostname | tr -d '\n' | tr -d '\r')

docker exec -it demo-stack-rabbit-2-1 rabbitmqctl stop_app
docker exec -it demo-stack-rabbit-2-1 rabbitmqctl reset
docker exec -it demo-stack-rabbit-2-1 rabbitmqctl join_cluster rabbit@$name
docker exec -it demo-stack-rabbit-2-1 rabbitmqctl start_app

docker exec -it demo-stack-rabbit-3-1 rabbitmqctl stop_app
docker exec -it demo-stack-rabbit-3-1 rabbitmqctl reset
docker exec -it demo-stack-rabbit-3-1 rabbitmqctl join_cluster rabbit@$name
docker exec -it demo-stack-rabbit-3-1 rabbitmqctl start_app
