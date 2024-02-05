#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it demo-stack-cockroach-1-1 ./cockroach init --host=cockroach-1:26357 --insecure
docker exec -i demo-stack-cockroach-1-1 ./cockroach sql --host=cockroach-1:26257  --insecure < init.sql
