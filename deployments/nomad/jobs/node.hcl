job "node" {
  type = "service"

  group "node" {
    count = 1

    network {
      mode = "bridge"

      port "node" {}
    }

    service {
      name     = "node"
      port     = "node"

      connect {
        sidecar_service {}
      }
    }

    task "node" {
      driver = "docker"

      config {
        image = "registry.com.gecko/node"
      }
    }
  }
}
