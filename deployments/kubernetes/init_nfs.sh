#!/bin/bash -ex

cd "$(dirname "$0")"

sudo apt install -y nfs-kernel-server
sudo mkdir -p /mnt/nfs_share
sudo chown -R nobody:nogroup /mnt/nfs_share
sudo echo /mnt/nfs_share *(rw,sync,no_subtree_check,no_root_squash,insecure) >> /etc/exports

sudo exportfs -a
sudo systemctl restart nfs-kernel-server
sudo exportfs -v

sudo mkdir -p /mnt/nfs_share/certs
sudo mkdir -p /mnt/nfs_share/influx
sudo mkdir -p /mnt/nfs_share/mongo
sudo mkdir -p /mnt/nfs_share/producer
sudo mkdir -p /mnt/nfs_share/rabbit-1
sudo mkdir -p /mnt/nfs_share/roach-certs
sudo mkdir -p /mnt/nfs_share/roach-1
