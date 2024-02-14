# Demo

Infrastructure and stack deployment.

## How to setup DNS

Add following entries to hosts file:

- 127.0.0.1 jenkins.com.gecko
- 127.0.0.1 registry.com.gecko
- 127.0.0.1 registry-ui.com.gecko
- 127.0.0.1 api.com.gecko
- 127.0.0.1 consul.com.gecko
- 127.0.0.1 grafana.com.gecko
- 127.0.0.1 nakama.com.gecko
- 127.0.0.1 receiver.com.gecko

## How to run CI

```bash
ssl/init_ssl.sh
ci/run.sh
```

## How to run demo on Docker Compose

```bash
stack/build.sh
stack/run.sh cockroach-1 cockroach-2 cockroach-3
stack/cockroach/init_ssl.sh
stack/run.sh
stack/cockroach/init_cluster.sh
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
- CockroachDB - <http://localhost:8080>
- Grafana - <https://grafana.com.gecko:4000> (admin, admin)
- HAProxy - <http://localhost:8404/stats>
- InfluxDB - <http://localhost:8086> (go_user, go_password)
- Mongo Express - <http://localhost:8881>
- Nakama - <https://nakam.com.gecko:4000> (admin, password)
- Prometheus - <http://localhost:9090>
- RabbitMQ - <http://localhost:15672> (guest, guest)
- RedisInsight - <http://localhost:8001>
