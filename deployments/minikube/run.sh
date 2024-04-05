#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl delete secret src-tls | true

kubectl create secret tls src-tls \
    --key ../../ssl/data/client-key.pem \
    --cert ../../ssl/data/client-chain.pem

find deployments -name "*.yaml" -exec kubectl apply --filename {} \;
find ingress     -name "*.yaml" -exec kubectl apply --filename {} \;
find pods        -name "*.yaml" -exec kubectl apply --filename {} \;
find services    -name "*.yaml" -exec kubectl apply --filename {} \;
find volumes     -name "*.yaml" -exec kubectl apply --filename {} \;
