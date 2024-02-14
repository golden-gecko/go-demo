job "mongo" {
  type = "service"

  group "mongo-group" {
    count = 1
    network {
      port "mongo-port" {
        to = 27017
      }
    }

    service {
      name     = "mongo-service"
      port     = "mongo-port"
      provider = "nomad"
    }

    task "node-task" {
      driver = "docker"

      config {
        image = "registry.com.gecko/mongo"
        ports = ["mongo-port"]
      }
    }
  }
}
