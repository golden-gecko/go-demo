#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it go-demo_cockroach_1_1 ./cockroach init --insecure

docker exec -it go-demo_cockroach_1_1 ./cockroach sql --insecure --execute "CREATE DATABASE go_demo"
docker exec -it go-demo_cockroach_1_1 ./cockroach sql --insecure --execute "CREATE USER go_user"
docker exec -it go-demo_cockroach_1_1 ./cockroach sql --insecure --execute "GRANT ALL ON DATABASE go_demo TO go_user WITH GRANT OPTION"
docker exec -it go-demo_cockroach_1_1 ./cockroach sql --insecure --execute "GRANT ALL ON TABLE go_demo.public.* TO go_user WITH GRANT OPTION"
