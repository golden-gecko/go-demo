job "node" {
  type = "service"

  group "node-group" {
    count = 1
    network {
      port "node-port" {
        to = 2000
      }
    }

    service {
      name     = "node-service"
      port     = "node-port"
      provider = "nomad"
    }

    task "node-task" {
      driver = "docker"

      config {
        image = "registry.com.gecko/node"
        ports = ["node-port"]
      }
    }
  }
}
