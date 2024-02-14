#!/bin/bash -ex

cd "$(dirname "$0")"

sudo nomad agent -dev \
    -bind 0.0.0.0 \
    -network-interface='{{ GetDefaultInterfaces | attr "name" }}'
