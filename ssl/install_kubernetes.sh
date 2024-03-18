#!/bin/bash -ex

cd "$(dirname "$0")"

sudo cp -r data/* /mnt/nfs_share/certs
