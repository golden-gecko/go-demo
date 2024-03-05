#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl apply --filename volumes/certs-claim.yaml
kubectl apply --filename volumes/certs-volume.yaml

kubectl apply --filename volumes/roach-1-claim.yaml
kubectl apply --filename volumes/roach-1-volume.yaml

kubectl apply --filename volumes/roach-certs-claim.yaml
kubectl apply --filename volumes/roach-certs-volume.yaml

kubectl apply --filename deployments/api.yaml
kubectl apply --filename deployments/node.yaml
# kubectl apply --filename deployments/roach.yaml

kubectl apply --filename pods/roach.yaml

kubectl apply --filename services/api.yaml
kubectl apply --filename services/node.yaml
# kubectl apply --filename services/roach.yaml
