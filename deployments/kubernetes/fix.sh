#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl -n kube-system rollout restart deployment coredns
