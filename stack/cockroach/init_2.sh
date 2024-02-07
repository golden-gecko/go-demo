#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it demo-stack-cockroach-1-1 ./cockroach init --certs-dir=/certs --host=cockroach-1:26357

docker exec -i -v $(pwd)/certs:/certs demo-stack-cockroach-1-1 ./cockroach sql --certs-dir=/certs --host=cockroach-1:26257 < init.sql
