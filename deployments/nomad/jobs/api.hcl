job "api" {
  type = "service"

  group "api-group" {
    count = 1
    network {
      port "api-port" {
        to = 6000
      }
    }

    service {
      name     = "api-service"
      port     = "api-port"
      provider = "nomad"
    }

    task "api-task" {
      driver = "docker"

      config {
        image = "registry.com.gecko/api"
        ports = ["api-port"]
      }
    }
  }
}
