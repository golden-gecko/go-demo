#!/bin/bash -ex

cd "$(dirname "$0")"

/nakama/nakama migrate up \
    --database.address 'root@roach-1:26257?sslmode=require&sslrootcert=/certs/ca.crt&sslcert=/certs/client.root.crt&sslkey=/certs/client.root.key' 

/nakama/nakama --name nakama1 \
    --database.address 'root@roach-1:26257?sslmode=require&sslrootcert=/certs/ca.crt&sslcert=/certs/client.root.crt&sslkey=/certs/client.root.key' \
    --logger.level DEBUG \
    --session.token_expiry_sec 7200 \
    --metrics.prometheus_port 9104
