#!/bin/bash -ex

cd "$(dirname "$0")"

echo MQATXXBUYSOQAYUNJPWZ > /var/lib/rabbitmq/.erlang.cookie
chmod 400 /var/lib/rabbitmq/.erlang.cookie

if [ -z "$1" ]
then
    rabbitmq-server
else
    # rabbitmqctl stop_app
    rabbitmqctl reset
    rabbitmqctl join_cluster rabbit@$1
    rabbitmqctl start_app    
fi

#cookie=$(docker exec -it go-demo-rabbit-1-1 sh -c "cat /var/lib/rabbitmq/.erlang.cookie")
#hostname=$(docker exec -it go-demo-rabbit-1-1 hostname)

#docker exec -it go-demo-rabbit-2-1 sh -c "echo > $cookie /var/lib/rabbitmq/.erlang.cookie"
#docker exec -it go-demo-rabbit-2-1 sh -c "chmod 400 /var/lib/rabbitmq/.erlang.cookie"

#docker exec -it go-demo-rabbit-2-1 rabbitmqctl stop_app
#docker exec -it go-demo-rabbit-2-1 rabbitmqctl reset
#docker exec -it go-demo-rabbit-2-1 rabbitmqctl join_cluster rabbit@$hostname
#docker exec -it go-demo-rabbit-2-1 rabbitmqctl start_app

#docker exec -it go-demo-rabbit-3-1 rabbitmqctl stop_app
#docker exec -it go-demo-rabbit-3-1 rabbitmqctl reset
#docker exec -it go-demo-rabbit-3-1 rabbitmqctl join_cluster rabbit@$hostname
#docker exec -it go-demo-rabbit-3-1 rabbitmqctl start_app
