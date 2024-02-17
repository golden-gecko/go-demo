# Demo

Infrastructure and stack deployment.

## How to setup DNS

Add following entries to hosts file and points them to localhost or Docker network IP:

```text
172.22.217.94 api.com.gecko
172.22.217.94 consul.com.gecko
172.22.217.94 grafana.com.gecko
172.22.217.94 haproxy.com.gecko
172.22.217.94 influx.com.gecko
172.22.217.94 jenkins.com.gecko
172.22.217.94 mongo.com.gecko
172.22.217.94 nakama.com.gecko
172.22.217.94 nomad.com.gecko
172.22.217.94 prometheus.com.gecko
172.22.217.94 rabbit.com.gecko
172.22.217.94 receiver.com.gecko
172.22.217.94 redis.com.gecko
172.22.217.94 registry.com.gecko
172.22.217.94 registry-ui.com.gecko
172.22.217.94 roach.com.gecko
```

## How to setup Docker

This demo is using self-signed certificates. Mark registry as insecure in Docker configuration.

```text
"insecure-registries": ["registry.com.gecko"]
```

## How to run CI

```bash
ssl/init_ssl.sh
ci/run.sh
```

## How to run demo on Docker Compose

```bash
stack/build.sh
stack/run.sh roach-1 roach-2 roach-3
stack/roach/init_ssl.sh
stack/run.sh
stack/roach/init_cluster.sh
stack/rabbit/init_cluster.sh
```

## How to run demo on Nomad

1. Install Nomad.

   <https://developer.hashicorp.com/nomad/tutorials/get-started/gs-install>

2. Create cluster.

   ```bash
   stack/build.sh
   deployments/nomad/init_cluster.sh
   ```

## UI

UI dashboards:

- Jenkins - <https://jenkins.com.gecko> (admin, ...)
- Docker Registry - <https://registry-ui.com.gecko>
- CockroachDB - <https://roach.com.gecko:4000> (go_user, go_password)
- Grafana - <https://grafana.com.gecko:4000> (admin, admin)
- HAProxy - <https://haproxy.com.gecko:4000/stats>
- InfluxDB - <https://influx.com.gecko:4000> (go_user, go_password)
- Mongo Express - <https://mongo.com.gecko:4000>
- Nakama - <https://nakama.com.gecko:4000> (admin, password)
- Prometheus - <https://prometheus.com.gecko:4000>
- RabbitMQ - <https://rabbit.com.gecko:4000> (guest, guest)
- RedisInsight - <https://redis.com.gecko:4000>
