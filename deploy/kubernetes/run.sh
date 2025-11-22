#!/bin/bash -ex

cd "$(dirname "$0")"

# kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
# kubectl -n kubernetes-dashboard apply -f admin-role.yml
# kubectl -n kubernetes-dashboard get secret admin-user-secret -o jsonpath="{.data.token}" | base64 -d

# find *.yml -print0 | xargs -n 1 -0 kubectl apply -f

# CockroachDB

kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/install/crds.yaml
kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/install/operator.yaml
kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/examples/example.yaml

kubectl create -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/examples/client-secure-operator.yaml

# InfluxDB

# Mongo

# Mongo Express

# RabbitMQ

# Redis

# Redis Insight

# Status
kubectl get pods --all-namespaces
kubectl get pv --all-namespaces
kubectl get service --all-namespaces
