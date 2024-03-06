#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl apply --filename volumes/certs-claim.yaml
kubectl apply --filename volumes/certs-volume.yaml
kubectl apply --filename volumes/influx-claim.yaml
kubectl apply --filename volumes/influx-volume.yaml
kubectl apply --filename volumes/mongo-claim.yaml
kubectl apply --filename volumes/mongo-volume.yaml
kubectl apply --filename volumes/roach-1-claim.yaml
kubectl apply --filename volumes/roach-1-volume.yaml
kubectl apply --filename volumes/roach-2-claim.yaml
kubectl apply --filename volumes/roach-2-volume.yaml
kubectl apply --filename volumes/roach-3-claim.yaml
kubectl apply --filename volumes/roach-3-volume.yaml
kubectl apply --filename volumes/roach-certs-claim.yaml
kubectl apply --filename volumes/roach-certs-volume.yaml

kubectl apply --filename deployments/api.yaml
kubectl apply --filename deployments/consumer.yaml
kubectl apply --filename deployments/node.yaml
kubectl apply --filename deployments/pistache.yaml
kubectl apply --filename deployments/python.yaml
kubectl apply --filename deployments/receiver.yaml

kubectl apply --filename pods/influx.yaml
kubectl apply --filename pods/mongo.yaml
kubectl apply --filename pods/producer.yaml
kubectl apply --filename pods/rabbit.yaml
kubectl apply --filename pods/roach.yaml

kubectl apply --filename services/api.yaml
kubectl apply --filename services/node.yaml
