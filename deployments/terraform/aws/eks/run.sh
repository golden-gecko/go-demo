#!/bin/bash -ex

cd "$(dirname "$0")"

# ssh-keygen -t rsa -b 4096 -m pem -f tutorial_kp
# openssl rsa -in tutorial_kp -outform pem

terraform init
terraform apply -auto-approve -var-file="secrets.tfvars"

aws eks update-kubeconfig --region eu-central-1 --name education-eks

# kubectl config use-context arn:aws:eks:eu-central-1:778117204733:cluster/education-eks
