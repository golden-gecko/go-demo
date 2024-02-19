job "mongo" {
  type = "service"

  group "mongo" {
    count = 1

    service {
      name     = "mongo"
      port     = 27017
    }

    task "node" {
      driver = "docker"

      config {
        image = "registry.com.gecko/mongo"
        ports = [27017]
      }
    }
  }
}
