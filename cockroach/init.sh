#!/bin/bash -ex

cd "$(dirname "$0")"

# docker volume rm cockroach-certs
# docker volume create cockroach-certs

docker run -it -v cockroach-certs:/certs bash -c "rm /certs/*"

docker run -it -v cockroach-certs:/certs cockroachdb/cockroach:v23.1.13 cert create-ca --certs-dir=/certs --ca-key=/certs/ca.key
docker run -it -v cockroach-certs:/certs cockroachdb/cockroach:v23.1.13 cert create-node localhost cockroach-1 --certs-dir=certs --ca-key=certs/ca.key
# docker run -it -v cockroach-certs:/certs cockroachdb/cockroach:v23.1.13 cert create-client root --certs-dir=certs --ca-key=certs/ca.key

# docker exec -it go-demo-cockroach-1-1 ./cockroach init --host=cockroach-1:26357
# docker exec -i go-demo-cockroach-1-1 ./cockroach sql --host=cockroach-1:26257 < init.sql
