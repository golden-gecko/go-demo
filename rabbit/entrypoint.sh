#!/bin/bash -ex

cd "$(dirname "$0")"

echo MQATXXBUYSOQAYUNJPWZ > /var/lib/rabbitmq/.erlang.cookie
chmod 400 /var/lib/rabbitmq/.erlang.cookie

rabbitmq-server start
