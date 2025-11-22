#!/bin/bash -ex

cd "$(dirname "$0")"

hostname=$(docker exec -it go-demo-rabbit-1-1 hostname)

docker exec -it go-demo-rabbit-2-1 rabbitmqctl stop_app
docker exec -it go-demo-rabbit-2-1 rabbitmqctl reset
docker exec -it go-demo-rabbit-2-1 rabbitmqctl join_cluster rabbit@$hostname
docker exec -it go-demo-rabbit-2-1 rabbitmqctl start_app

docker exec -it go-demo-rabbit-3-1 rabbitmqctl stop_app
docker exec -it go-demo-rabbit-3-1 rabbitmqctl reset
docker exec -it go-demo-rabbit-3-1 rabbitmqctl join_cluster rabbit@$hostname
docker exec -it go-demo-rabbit-3-1 rabbitmqctl start_app
