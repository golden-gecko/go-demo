#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl apply --filename volumes/roach-claim.yaml
kubectl apply --filename volumes/roach.yaml

kubectl apply --filename deployments/api.yaml
kubectl apply --filename deployments/node.yaml
kubectl apply --filename deployments/roach.yaml

kubectl apply --filename services/api.yaml
kubectl apply --filename services/node.yaml
kubectl apply --filename services/roach.yaml
