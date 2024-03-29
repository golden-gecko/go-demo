#!/bin/bash -ex

cd "$(dirname "$0")"

terraform init
terraform apply -auto-approve

# aws eks --region $(terraform output -raw region) update-kubeconfig --name $(terraform output -raw cluster_name)
# kubectl config use-context arn:aws:eks:eu-central-1:778117204733:cluster/education-eks
