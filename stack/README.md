# Stack

Stack for gathering and processing data.

## How to setup DNS

Add following domain to DNS and point them to Docker IP.

- api.com.gecko
- consul.com.gecko
- grafana.com.gecko
- nakama.com.gecko
- receiver.com.gecko

## How to run

```bash
./run.sh
```

## How to stop

```bash
./stop.sh
```

## UI

- CockroachDB - <http://localhost:8080>
- Grafana - <https://grafana.com.gecko:2000> (admin, admin)
- HAProxy - <http://localhost:8404/stats>
- InfluxDB - <http://localhost:8086> (go_user, go_password)
- Mongo Express - <http://localhost:8881>
- Nakama - <http://localhost:7351> (admin, password)
- Prometheus - <http://localhost:9090>
- RabbitMQ - <http://localhost:15672> (guest, guest)
- RedisInsight - <http://localhost:8001>
