#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it go-demo-cockroach_1-1 ./cockroach init --host=cockroach_1:26357 --insecure

docker exec -i go-demo-cockroach_1-1 ./cockroach sql --host=cockroach_1:26257 --insecure < init.sql
