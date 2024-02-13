#!/bin/bash -ex

cd "$(dirname "$0")"

docker run -it -v demo-certs:/certs -v $(pwd):/app busybox bash /app/init_ssl_docker.sh
