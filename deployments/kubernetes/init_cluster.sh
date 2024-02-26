#!/bin/bash -ex

cd "$(dirname "$0")"

minikube start

minikube addons enable ingress
minikube addons enable metrics-server

minikube cp /usr/local/share/ca-certificates /usr/local/share/ca-certificates
minikube ssh "sudo update-ca-certificates --fresh"
