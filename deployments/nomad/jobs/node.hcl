job "node" {
  type = "service"

  group "node" {
    count = 1

    network {
      mode = "bridge"
    }

    service {
      name     = "node"
      port     = 2000

      connect {
        sidecar_service {}
      }
    }

    task "node" {
      driver = "docker"

      config {
        image = "registry.com.gecko/node"
        ports = ["node"]
      }
    }
  }
}
