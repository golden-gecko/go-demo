#!/bin/bash -ex

cd "$(dirname "$0")"

# rabbit
rabbit_name_1=$(docker compose --project-name demo-src ps --quiet rabbit-1)
rabbit_name_2=$(docker compose --project-name demo-src ps --quiet rabbit-2)
rabbit_name_3=$(docker compose --project-name demo-src ps --quiet rabbit-3)

name=$(docker exec -it $rabbit_name_1 hostname | tr -d '\n' | tr -d '\r')

docker exec -it $rabbit_name_2 rabbitmqctl stop_app
docker exec -it $rabbit_name_2 rabbitmqctl reset
docker exec -it $rabbit_name_2 rabbitmqctl join_cluster rabbit@$name
docker exec -it $rabbit_name_2 rabbitmqctl start_app

docker exec -it $rabbit_name_3 rabbitmqctl stop_app
docker exec -it $rabbit_name_3 rabbitmqctl reset
docker exec -it $rabbit_name_3 rabbitmqctl join_cluster rabbit@$name
docker exec -it $rabbit_name_3 rabbitmqctl start_app

# roach
roach_name=$(docker compose --project-name demo-src ps --quiet roach-1)

docker exec -it $roach_name sh -c "cockroach init --certs-dir=/certs/node-1 --host=roach-1:26357"
docker exec -i $roach_name sh -c "cockroach sql --certs-dir=/certs/node-1 --host=roach-1:26257" < ../../src/roach/init.sql
