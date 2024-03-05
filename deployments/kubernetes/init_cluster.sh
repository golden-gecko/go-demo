#!/bin/bash -ex

minikube start --mount --mount-string="$1:/opt/data"

minikube addons enable ingress
minikube addons enable metrics-server

minikube ssh "sudo rm -f /usr/local/share/ca-certificates/gecko.crt"
minikube ssh "sudo ln -s /opt/data/certs/ca-cert.pem /usr/local/share/ca-certificates/gecko.crt"
minikube ssh "sudo update-ca-certificates"
