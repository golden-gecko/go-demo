#!/bin/bash -ex

cd "$(dirname "$0")"

ssh-keygen -t rsa -b 4096 -m pem -f go-demo && openssl rsa -in go-demo -outform pem

terraform init
terraform apply -auto-approve -var-file="variables.tfvars"
