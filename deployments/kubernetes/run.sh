#!/bin/bash -ex

cd "$(dirname "$0")"

find deployments -name "*.yaml" -exec kubectl apply --filename {} \;
find pods -name "*.yaml" -exec kubectl apply --filename {} \;
find services -name "*.yaml" -exec kubectl apply --filename {} \;
find volumes -name "*.yaml" -exec kubectl apply --filename {} \;
