# Demo

Infrastructure and stack deployment.

## How to setup DNS

Add following entries to hosts file and points them to localhost or Docker network IP:

```bash
172.22.217.94 api.com.gecko
172.22.217.94 consul.com.gecko
172.22.217.94 grafana.com.gecko
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

- Jenkins - <https://jenkins.com.gecko>
- Docker Registry - <https://registry-ui.com.gecko>
- CockroachDB - <https://roach.com.gecko:4000>
- Grafana - <https://grafana.com.gecko:4000> (admin, admin)

HAProxy - <http://localhost:8404/stats>
InfluxDB - <http://localhost:8086> (go_user, go_password)

- Mongo Express - <https://mongo.com.gecko:4000>
- Nakama - <https://nakama.com.gecko:4000> (admin, password)
- Prometheus - <https://prometheus.com.gecko:4000>
- RabbitMQ - <https://rabbit.com.gecko:4000> (guest, guest)
- RedisInsight - <https://redis.com.gecko:4000>
