#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it demo-stack-cockroach-1-1 ./cockroach init --certs-dir=/certs/node-1 --host=cockroach-1:26357
docker exec -i demo-stack-cockroach-1-1 ./cockroach sql --certs-dir=/certs/node-1 --host=cockroach-1:26257 < init.sql
