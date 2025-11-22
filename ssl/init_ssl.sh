#!/bin/bash -ex

cd "$(dirname "$0")"

docker build -t demo-tools .
docker run -i --name demo-certs -v demo-certs:/certs -v $(pwd):/app -t demo-tools bash /app/init_ssl_docker.sh
docker rm demo-certs
