#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl delete secret stack-tls | true

find deployments -name "*.yaml" | xargs -I{} basename {} ".yaml" | xargs kubectl delete --ignore-not-found=true deployment
find ingress     -name "*.yaml" | xargs -I{} basename {} ".yaml" | xargs kubectl delete --ignore-not-found=true ingress
find pods        -name "*.yaml" | xargs -I{} basename {} ".yaml" | xargs kubectl delete --ignore-not-found=true pod
find services    -name "*.yaml" | xargs -I{} basename {} ".yaml" | xargs kubectl delete --ignore-not-found=true service
find volumes     -name "*.yaml" | xargs -I{} basename {} ".yaml" | xargs kubectl delete --ignore-not-found=true pvc
find volumes     -name "*.yaml" | xargs -I{} basename {} ".yaml" | xargs kubectl delete --ignore-not-found=true pv
