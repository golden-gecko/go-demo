#!/bin/bash -ex

cd "$(dirname "$0")"

docker rm demo-roach-certs || true
docker run -i --name demo-roach-certs -v $(pwd)/data:/certs -v $(pwd):/app -t registry.com.gecko/roach bash /app/generate.sh /certs
docker rm demo-roach-certs
