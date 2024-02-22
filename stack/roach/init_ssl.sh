#!/bin/bash -ex

cd "$(dirname "$0")"

docker run -i --name demo-roach-certs -v demo-roach-certs:/certs -v $(pwd):/app -t registry.com.gecko/roach bash /app/init_ssl_docker.sh
docker rm demo-roach-certs
