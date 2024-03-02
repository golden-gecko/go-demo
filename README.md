# Demo

Infrastructure and stack deployment.

## How to setup DNS

Add following entries to hosts file and points them to localhost or Docker network IP:

```text
127.0.0.1 api.com.gecko
127.0.0.1 consul.com.gecko
127.0.0.1 grafana.com.gecko
127.0.0.1 haproxy.com.gecko
127.0.0.1 influx.com.gecko
127.0.0.1 jenkins.com.gecko
127.0.0.1 mongo.com.gecko
127.0.0.1 nakama.com.gecko
127.0.0.1 nomad.com.gecko
127.0.0.1 portainer.com.gecko
127.0.0.1 prometheus.com.gecko
127.0.0.1 rabbit.com.gecko
127.0.0.1 receiver.com.gecko
127.0.0.1 redis.com.gecko
127.0.0.1 registry.com.gecko
127.0.0.1 registry-ui.com.gecko
127.0.0.1 roach.com.gecko
```

## How to setup SSL

1. Generate certificates.

   ```bash
   ssl/init_ssl.sh data/certs
   ```

2. Install CA certificate.

   ```bash
   ssl/install.sh data/certs
   ```

3. Create docker volume with certificates.

   ```bash
   ssl/init_docker.sh
   ```

## How to run CI

```bash
ci/run.sh
```

## How to run on Docker Compose

```bash
deployments/compose/build.sh
stack/roach/init_ssl.sh
deployments/compose/run.sh
stack/roach/init_cluster.sh
stack/rabbit/init_cluster.sh
```

## How to run on Kubernetes

1. Install Kubernetes.

2. Create cluster.

   ```bash
   minikube start
   ```

## How to run on Nomad

1. Install Nomad.

   <https://developer.hashicorp.com/nomad/tutorials/get-started/gs-install>

2. Create cluster.

   ```bash
   deployments/compose/build.sh
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
- Portainer <https://portainer.com.gecko> (admin, ...)
- Prometheus - <https://prometheus.com.gecko:4000>
- RabbitMQ - <https://rabbit.com.gecko:4000> (guest, guest)
- RedisInsight - <https://redis.com.gecko:4000>

## Bugs

- RedisInsight is not working.
