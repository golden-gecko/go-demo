#!/bin/bash -ex

cd "$(dirname "$0")"

# API
kubectl delete -f api/api.yaml | true

# CockroachDB
kubectl delete -f cockroachdb/crds.yaml | true
kubectl delete -f cockroachdb/operator.yaml | true
kubectl delete -f cockroachdb/example.yaml | true
