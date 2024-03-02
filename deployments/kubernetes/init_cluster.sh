#!/bin/bash -ex

cd "$(dirname "$0")"

minikube start --mount --mount-string="../../data:/mnt/data"

minikube addons enable ingress
minikube addons enable metrics-server

minikube cp /mnt/data/certs/ca-cert.pem /usr/local/share/ca-certificates/gecko.crt
minikube ssh "sudo update-ca-certificates"
