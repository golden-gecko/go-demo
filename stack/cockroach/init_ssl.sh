#!/bin/bash -ex

cd "$(dirname "$0")"

MSYS_NO_PATHCONV=1 # for Bash on Windows

mkdir -p certs
rm -fr certs/*.crt certs/*.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-ca \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

mkdir -p certs/node-1 certs/node-2 certs/node-3
rm -fr certs/node-*/*.crt certs/node-*/*.key

cp certs/ca.crt certs/node-1/
cp certs/ca.crt certs/node-2/
cp certs/ca.crt certs/node-3/

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-node cockroach-1 \
    --certs-dir=/certs/node-1 \
    --ca-key=/certs/ca.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-node cockroach-2 \
    --certs-dir=/certs/node-2 \
    --ca-key=/certs/ca.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-node cockroach-3 \
    --certs-dir=/certs/node-3 \
    --ca-key=/certs/ca.key

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-client root \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

cp certs/client.root.crt certs/node-1/
cp certs/client.root.crt certs/node-2/
cp certs/client.root.crt certs/node-3/

cp certs/client.root.key certs/node-1/
cp certs/client.root.key certs/node-2/
cp certs/client.root.key certs/node-3/

docker run -it -v $(pwd)/certs:/certs registry.com.gecko/cockroach cert create-client go_user \
    --certs-dir=/certs \
    --ca-key=/certs/ca.key

mkdir -p ../services/services/api/certs
rm -fr ../services/services/api/certs/*.crt ../services/services/api/certs/*.key

cp certs/ca.crt ../services/services/api/certs/
cp certs/client.go_user.crt ../services/services/api/certs/
cp certs/client.go_user.key ../services/services/api/certs/
