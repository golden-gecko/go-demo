#!/bin/bash -ex

cd "$(dirname "$0")"

# API
kubectl apply -f api/api.yaml

# CockroachDB
kubectl apply -f cockroachdb/crds.yaml
kubectl apply -f cockroachdb/operator.yaml
kubectl apply -f cockroachdb/example.yaml

# CockroachDB client
kubectl create -f cockroachdb/client-secure-operator.yaml

# CockroachDB forward
# kubectl port-forward service/cockroachdb 8080:8080

# InfluxDB

# Mongo

# Mongo Express

# RabbitMQ

# Redis

# Redis Insight

# Status
# kubectl get deployment --all-namespaces
# kubectl get pod --all-namespaces
# kubectl get pv --all-namespaces
# kubectl get pvc --all-namespaces
# kubectl get secret --all-namespaces
# kubectl get service --all-namespaces

# Logs
# kubectl logs pod_name
