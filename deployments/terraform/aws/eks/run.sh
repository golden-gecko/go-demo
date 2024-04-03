#!/bin/bash -ex

cd "$(dirname "$0")"

terraform init
terraform apply -auto-approve

<<<<<<< Updated upstream
aws eks --region $(terraform output -raw region) update-kubeconfig --name $(terraform output -raw cluster_name)
kubectl config use-context arn:aws:eks:$(terraform output -raw region):$(terraform output -raw account_id):cluster/$(terraform output -raw cluster_name)
=======
region=$(terraform output -raw region)
cluster_name=$(terraform output -raw cluster_name)

# aws ecr get-login-password --region $region | docker login --username AWS --password-stdin 778117204733.dkr.ecr.$region.amazonaws.com

# aws eks --region $region update-kubeconfig --name $cluster_name

# kubectl config use-context arn:aws:eks:eu-central-1:778117204733:cluster/$(cluster_name)
>>>>>>> Stashed changes
