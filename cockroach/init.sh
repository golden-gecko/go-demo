#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it go-demo_cockroach_1_1 ./cockroach init --host=cockroach_1:26357 --insecure

docker exec -i go-demo_cockroach_1_1 ./cockroach sql --host=cockroach_1:26257 --insecure < init.sql
