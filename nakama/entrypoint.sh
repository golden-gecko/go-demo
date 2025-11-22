#!/bin/bash -ex

cd "$(dirname "$0")"

/nakama/nakama migrate up --database.address root@cockroach-1:26257
/nakama/nakama --name nakama1 --database.address root@cockroach-1:26257 --logger.level DEBUG --session.token_expiry_sec 7200 --metrics.prometheus_port 9104
