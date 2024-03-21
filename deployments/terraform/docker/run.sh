#!/bin/bash -ex

cd "$(dirname "$0")"

terraform init
terraform apply -auto-approve
