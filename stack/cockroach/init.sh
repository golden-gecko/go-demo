#!/bin/bash -ex

cd "$(dirname "$0")"

mkdir -p certs
rm -fr certs/*.crt certs/*.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-ca \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-node localhost $(hostname) \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-client root \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

# docker exec -it demo-stack-cockroach-1-1 ./cockroach init --certs-dir=/certs --host=cockroach-1:26357
# docker exec -i demo-stack-cockroach-1-1 ./cockroach sql --certs-dir=/certs --host=cockroach-1:26257 < init.sql
