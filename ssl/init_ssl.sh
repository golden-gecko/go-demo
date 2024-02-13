#!/bin/bash -ex

cd "$(dirname "$0")"

docker build -t demo-tools .
docker run -it -v demo-certs:/certs -v $(pwd):/app demo-tools bash /app/init_ssl_docker.sh
