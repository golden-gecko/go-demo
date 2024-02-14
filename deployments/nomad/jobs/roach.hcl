job "roach" {
  type = "service"

  group "roach-group" {
    count = 1

    network {
      port "roach-port" {
        to = 26257
      }
    }

    service {
      name     = "roach-service"
      port     = "roach-port"
      provider = "nomad"
    }

    task "roach-task" {
      driver = "docker"

      config {
        image   = "registry.com.gecko/roach"
        ports   = ["roach-port"]
        command = "start-single-node"
        args    = ["--insecure"]
      }

      resources {
        cpu    = 100
        memory = 1024
      }
    }
  }
}
