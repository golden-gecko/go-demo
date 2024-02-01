#!/bin/bash -ex

cd "$(dirname "$0")"

# API
kubectl delete -f api/api.yaml

# CockroachDB
kubectl delete -f cockroachdb/crds.yaml
kubectl delete -f cockroachdb/operator.yaml
kubectl delete -f cockroachdb/example.yaml
