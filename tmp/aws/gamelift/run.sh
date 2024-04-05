#!/bin/bash -ex

cd "$(dirname "$0")"

# ssh-keygen -t rsa -b 4096 -m pem -f tutorial_gamelift_kp
# openssl rsa -in tutorial_gamelift_kp -outform pem

terraform init
terraform apply -auto-approve
