#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl exec -i cockroachdb-client-secure \
    -- ./cockroach sql \
    --certs-dir=/cockroach/cockroach-certs \
    --host=cockroachdb-public < ..\..\cockroach\init.sql
