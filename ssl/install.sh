#!/bin/bash -ex

cd "$(dirname "$0")"

sudo cp data/certs/ca-cert.pem /usr/local/share/ca-certificates/gecko.crt
sudo update-ca-certificates
