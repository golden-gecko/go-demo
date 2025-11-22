# Go Demo

## Usage

### How to run

```bash
./run.sh
```

### How to deploy locally

```bash
./run.sh
```

### How to deploy using Jenkins

Run CI.

```bash
deploy/ci/run.sh
```

Generate SSH key.

```bash
deploy/ci/init.sh
```

Open Jenkins http://localhost:8888/.

Add SSH key to Jenkins and Git repository.

### How to deploy using Kubernetes

```bash
```

### UI

- Chronograf - http://localhost:8888/
- CockroachDB - http://localhost:8080/
- Grafana - http://localhost:3000/ (admin, admin)
- HAProxy - http://localhost:8404/stats
- InfluxDB - http://localhost:8086/ (go_user, go_password)
- Mongo Express - http://localhost:8881/
- Nakama - http://localhost:7351/ (admin, password)
- Prometheus - http://localhost:9090/
- RabbitMQ - http://localhost:15672/ (guest, guest)
- RedisInsight - http://localhost:8001/
