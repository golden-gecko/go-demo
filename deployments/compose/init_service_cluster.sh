#!/bin/bash -ex

cd "$(dirname "$0")"

# rabbit
name=$(docker exec -it demo-stack_rabbit-1_1 hostname | tr -d '\n' | tr -d '\r')

docker exec -it demo-stack_rabbit-2_1 rabbitmqctl stop_app
docker exec -it demo-stack_rabbit-2_1 rabbitmqctl reset
docker exec -it demo-stack_rabbit-2_1 rabbitmqctl join_cluster rabbit@$name
docker exec -it demo-stack_rabbit-2_1 rabbitmqctl start_app

docker exec -it demo-stack_rabbit-3_1 rabbitmqctl stop_app
docker exec -it demo-stack_rabbit-3_1 rabbitmqctl reset
docker exec -it demo-stack_rabbit-3_1 rabbitmqctl join_cluster rabbit@$name
docker exec -it demo-stack_rabbit-3_1 rabbitmqctl start_app

# roach
# docker exec -i -t demo-stack_roach-1_1 cockroach init --certs-dir=/certs/node-1 --host=roach-1:26357
# docker exec -i demo-stack_roach-1_1 cockroach sql --certs-dir=/certs/node-1 --host=roach-1:26257 < ../../stack/roach/init.sql
