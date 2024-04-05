#!/bin/bash -ex

cd "$(dirname "$0")"

# rabbit
# name=$(kubectl exec --stdin --tty -- rabbit-1 hostname | tr -d '\n' | tr -d '\r')

# kubectl exec --stdin --tty -- rabbit-2 rabbitmqctl stop_app
# kubectl exec --stdin --tty -- rabbit-2 rabbitmqctl reset
# kubectl exec --stdin --tty -- rabbit-2 rabbitmqctl join_cluster rabbit@$name
# kubectl exec --stdin --tty -- rabbit-2 rabbitmqctl start_app

# kubectl exec --stdin --tty -- rabbit-3 rabbitmqctl stop_app
# kubectl exec --stdin --tty -- rabbit-3 rabbitmqctl reset
# kubectl exec --stdin --tty -- rabbit-3 rabbitmqctl join_cluster rabbit@$name
# kubectl exec --stdin --tty -- rabbit-3 rabbitmqctl start_app

# roach
# kubectl exec --stdin --tty roach-1 -- cockroach init --certs-dir=/certs/node-1 --host=roach-1:26357
kubectl exec --stdin roach-1 -- cockroach sql --certs-dir=/certs/node-1 --host=roach-1:26257 < ../../src/roach/init.sql
