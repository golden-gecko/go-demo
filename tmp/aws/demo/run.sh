#!/bin/bash -ex

cd "$(dirname "$0")"

# ssh-keygen -t rsa -b 4096 -m pem -f tutorial_kp
# openssl rsa -in tutorial_kp -outform pem

terraform init
terraform apply -auto-approve -var-file="secrets.tfvars"
