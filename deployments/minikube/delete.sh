#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl delete pvc certs-claim
kubectl delete pv certs-volume
kubectl delete pvc influx-claim
kubectl delete pv influx-volume
kubectl delete pvc mongo-claim
kubectl delete pv mongo-volume
kubectl delete pvc rabbit-1-claim
kubectl delete pv rabbit-1-volume
kubectl delete pvc roach-1-claim
kubectl delete pv roach-1-volume
kubectl delete pvc roach-2-claim
kubectl delete pv roach-2-volume
kubectl delete pvc roach-3-claim
kubectl delete pv roach-3-volume
kubectl delete pvc roach-certs-claim
kubectl delete pv roach-certs-volume
