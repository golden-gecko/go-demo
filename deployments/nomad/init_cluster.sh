#!/bin/bash -ex

cd "$(dirname "$0")"

sudo mkdir -p \
    /opt/nomad/data/certs \
    /opt/nomad/data/roach-certs \
    /opt/nomad/data/roach-1 \
    /opt/nomad/data/roach-2 \
    /opt/nomad/data/roach-3

sudo nomad agent -dev-connect \
    -config=config.hcl \
    -network-interface='{{ GetDefaultInterfaces | attr "name" }}'
