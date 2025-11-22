#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -i go-demo_cockroach-1_1 ./cockroach sql --host=cockroach-1:26257 --insecure < "$@"
