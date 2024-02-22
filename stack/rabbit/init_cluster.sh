#!/bin/bash -ex

cd "$(dirname "$0")"

name=$(docker exec -it demo-stack_rabbit-1_1 hostname | tr -d '\n' | tr -d '\r')

docker exec -it demo-stack_rabbit-2_1 rabbitmqctl stop_app
docker exec -it demo-stack_rabbit-2_1 rabbitmqctl reset
docker exec -it demo-stack_rabbit-2_1 rabbitmqctl join_cluster rabbit@$name
docker exec -it demo-stack_rabbit-2_1 rabbitmqctl start_app

docker exec -it demo-stack_rabbit-3_1 rabbitmqctl stop_app
docker exec -it demo-stack_rabbit-3_1 rabbitmqctl reset
docker exec -it demo-stack_rabbit-3_1 rabbitmqctl join_cluster rabbit@$name
docker exec -it demo-stack_rabbit-3_1 rabbitmqctl start_app
