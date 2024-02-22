#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -i -t demo-stack_roach-1_1 cockroach init --certs-dir=/certs/node-1 --host=roach-1:26357
docker exec -i demo-stack_roach-1_1 cockroach sql --certs-dir=/certs/node-1 --host=roach-1:26257 < init.sql
