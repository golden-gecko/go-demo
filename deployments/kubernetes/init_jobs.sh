#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl apply --filename deployments/api.yaml
kubectl apply --filename deployments/node.yaml
kubectl apply --filename deployments/receiver.yaml

kubectl expose deployment/node-deployment --type="NodePort" --port 2000
