#!/bin/bash -ex

cd "$(dirname "$0")"

mkdir -p certs
rm -fr certs/*.crt certs/*.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-ca \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-node cockroach-1 \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-client root \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key
