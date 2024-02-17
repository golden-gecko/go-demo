job "mongo" {
  type = "service"

  group "mongo" {
    count = 1

    network {
      mode = "host"

      port "mongo" {
        to = 27017
      }
    }

    service {
      name     = "mongo"
      port     = "mongo"
      provider = "nomad"
    }

    task "node" {
      driver = "docker"

      config {
        image = "registry.com.gecko/mongo"
        ports = ["mongo"]
      }
    }
  }
}
