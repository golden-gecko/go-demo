#!/bin/bash -ex

sudo cp $1/ca-cert.pem /usr/local/share/ca-certificates/gecko.crt
sudo update-ca-certificates
