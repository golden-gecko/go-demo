#!/bin/bash -ex

cd "$(dirname "$0")"

docker run -it -v demo-stack_cockroach-certs:/certs -v $(pwd):/app registry.com.gecko/cockroach bash /app/init_ssl_docker.sh
