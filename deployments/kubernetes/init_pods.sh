#!/bin/bash -ex

cd "$(dirname "$0")"

kubectl exec --stdin --tty roach-deployment-779894f9cd-bqtvd -- ls
