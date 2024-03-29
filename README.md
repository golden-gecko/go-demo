# Demo

Infrastructure and stack deployment.

## How to setup DNS

Add following entries to hosts file and points them to localhost or Docker network IP:

```text
127.0.0.1 jenkins.com.gecko
127.0.0.1 nfs.com.gecko
127.0.0.1 portainer.com.gecko
127.0.0.1 registry.com.gecko
127.0.0.1 registry-ui.com.gecko

127.0.0.1 api.com.gecko
127.0.0.1 consul.com.gecko
127.0.0.1 grafana.com.gecko
127.0.0.1 haproxy.com.gecko
127.0.0.1 influx.com.gecko
127.0.0.1 mongo.com.gecko
127.0.0.1 nakama.com.gecko
127.0.0.1 rabbit.com.gecko
127.0.0.1 receiver.com.gecko
127.0.0.1 redis.com.gecko
127.0.0.1 roach.com.gecko
```

## How to setup SSL

1. Generate certificates.

   ```bash
   ssl/generate.sh
   ```

2. Install CA certificate in OS.

   ```bash
   ssl/install_local.sh
   ```

3. Restart docker.

   ```bash
   sudo systemctl restart docker
   ```

## How to run CI

1. Create docker volume with certificates.

   ```bash
   ssl/install_docker.sh
   ```

2. Run CI.

   ```bash
   ci/run.sh
   ```

## How to run on Docker Compose

```bash
deployments/compose/build.sh
stack/roach/ssl/generate_docker.sh
deployments/compose/run.sh
deployments/compose/init_service_cluster.sh
```

## How to run on Kubernetes (minikube)

1. Install minikube nad krew.

   - <https://minikube.sigs.k8s.io/docs/start/>
   - <https://krew.sigs.k8s.io/docs/user-guide/setup/install/>

   Run:

   ```bash
   kubectl krew install rabbitmq
   kubectl rabbitmq install-cluster-operator
   ```

2. Run.

   ```bash
   stack/roach/ssl/generate_local.sh
   stack/roach/ssl/install_kubernetes.sh
   deployments/kubernetes/run_cluster.sh
   deployments/compose/build.sh
   deployments/kubernetes/run.sh
   ```

## How to run on Kubernetes (EKS)

1. Run:

   ```bash
   deployments/aws/eks/run.sh
   ```

## How to run on Nomad (in progress)

1. Install Nomad.

   <https://developer.hashicorp.com/nomad/tutorials/get-started/gs-install>

2. Create cluster.

   ```bash
   deployments/compose/build.sh
   deployments/nomad/init_cluster.sh
   ```

## UI (valid with Docker Compose only)

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

## Roadmap

- Add CockroachSQL to haproxy.

  ```bash
  docker exec -it $roach_name sh -c "cockroach gen haproxy --certs-dir=/certs/node-1 --host=roach-1:26357 && cat haproxy.cfg" > ../../stack/haproxy/haproxy_roach.cfg
  ```
