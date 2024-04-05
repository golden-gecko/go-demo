#!/bin/bash -ex

cd "$(dirname "$0")"

terraform destroy -auto-approve
