#!/bin/bash -ex

cd "$(dirname "$0")"

curl -o cockroachdb/crds.yaml https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/install/crds.yaml
curl -o cockroachdb/operator.yaml https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/install/operator.yaml
curl -o cockroachdb/example.yaml https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/examples/example.yaml
curl -o cockroachdb/client-secure-operator.yaml https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.12.0/examples/client-secure-operator.yaml
