#!/bin/bash -ex

cd "$(dirname "$0")"

cd /certs

rm -fr *.crt *.key

cockroach cert create-ca \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

mkdir -p node-1 node-2 node-3
rm -fr node-*/*.crt node-*/*.key

cp ca.crt node-1/
cp ca.crt node-2/
cp ca.crt node-3/

cockroach cert create-node cockroach-1 \
    --certs-dir=/certs/node-1 \
    --ca-key=/certs/ca.key

cockroach cert create-node cockroach-2 \
    --certs-dir=/certs/node-2 \
    --ca-key=/certs/ca.key

cockroach cert create-node cockroach-3 \
    --certs-dir=/certs/node-3 \
    --ca-key=/certs/ca.key

cockroach cert create-client root \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

cp client.root.crt node-1/
cp client.root.crt node-2/
cp client.root.crt node-3/

cp client.root.key node-1/
cp client.root.key node-2/
cp client.root.key node-3/

cockroach cert create-client go_user \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key
