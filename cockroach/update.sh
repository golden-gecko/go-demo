#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -i go-demo_cockroach_1_1 ./cockroach sql --host=cockroach_1:26257 --insecure < "$@"
