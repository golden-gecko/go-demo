#!/bin/bash -ex

cd "$(dirname "$0")"

minikube start --nodes=3

minikube addons enable ingress
minikube addons enable metrics-server

minikube cp ../../ssl/data/certs/ca-cert.pem minikube:/usr/local/share/ca-certificates/gecko.crt
minikube cp ../../ssl/data/certs/ca-cert.pem minikube-m02:/usr/local/share/ca-certificates/gecko.crt
minikube cp ../../ssl/data/certs/ca-cert.pem minikube-m03:/usr/local/share/ca-certificates/gecko.crt

minikube ssh --node minikube "sudo update-ca-certificates"
minikube ssh --node minikube-m02 "sudo update-ca-certificates"
minikube ssh --node minikube-m03 "sudo update-ca-certificates"

minikube ssh --node minikube "sudo systemctl restart docker"
minikube ssh --node minikube-m02 "sudo systemctl restart docker"
minikube ssh --node minikube-m03 "sudo systemctl restart docker"
