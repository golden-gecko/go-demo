#!/bin/bash -ex

cd "$(dirname "$0")"

sudo cp /var/lib/docker/volumes/demo-certs/_data/ca-cert.pem /usr/local/share/ca-certificates/gecko.crt
sudo update-ca-certificates --fresh

sudo mkdir -p /etc/docker/certs.d/registry.com.gecko
sudo cp /var/lib/docker/volumes/demo-certs/_data/ca-cert.pem /etc/docker/certs.d/registry.com.gecko/ca.crt
