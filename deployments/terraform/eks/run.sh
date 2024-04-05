#!/bin/bash -ex

cd "$(dirname "$0")"

terraform init
terraform apply -auto-approve

account_id=$(terraform output -raw account_id)
region=$(terraform output -raw region)
cluster_name=$(terraform output -raw cluster_name)

aws ecr get-login-password --region ${region} | docker login --username AWS --password-stdin ${account_id}.dkr.ecr.${region}.amazonaws.com
aws eks --region ${region} update-kubeconfig --name ${cluster_name}
kubectl config use-context arn:aws:eks:${region}:${account_id}:cluster/${cluster_name}
