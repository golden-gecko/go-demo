#!/bin/bash -ex

cd "$(dirname "$0")"

docker build -t demo-tools .
docker stop demo-certs || true
docker rm demo-certs || true
docker create --name demo-certs -v demo-certs:/certs -v $(pwd):/app -t demo-tools bash
docker cp data/. demo-certs:/certs
docker stop demo-certs
docker rm demo-certs
