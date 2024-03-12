#!/bin/bash -ex

cd "$(dirname "$0")"

minikube start --nodes=3

minikube addons enable ingress
minikube addons enable metrics-server

declare -a nodes=("minikube" "minikube-m02" "minikube-m03")

for i in "${nodes[@]}"
do
    minikube cp ../../ssl/data/ca-cert.pem $i:/usr/local/share/ca-certificates/gecko.crt

    minikube ssh --node $i "sudo update-ca-certificates"
    minikube ssh --node $i "sudo systemctl restart docker"
done
