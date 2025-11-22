#!/bin/bash -ex

cd "$(dirname "$0")"

if [ -z "${MASTER_HOST}"] && [ -z "${MASTER_PORT}" ]
then
    redis-server
else
    redis-server --slaveof ${MASTER_HOST} ${MASTER_PORT}
fi
