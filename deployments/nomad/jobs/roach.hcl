job "roach" {
  type = "service"

  group "roach" {
    count = 1

    network {
      mode = "host"

      port "roach" {
        to = 26257
      }
    }

    service {
      name     = "roach"
      port     = "roach"
      provider = "nomad"
    }

    volume "roach-certs" {
      type      = "host"
      read_only = true
      source    = "roach-certs"
    }

    volume "roach-1" {
      type      = "host"
      read_only = false
      source    = "roach-1"
    }

    task "roach" {
      driver = "docker"

      config {
        image   = "registry.com.gecko/roach"
        ports   = ["roach"]
        command = "start-single-node"
        args    = ["--certs-dir=/certs/node-1"]
      }

      resources {
        cpu    = 100
        memory = 1024
      }

      volume_mount {
        volume      = "roach-certs"
        destination = "/certs"
        read_only   = true
      }

      volume_mount {
        volume      = "roach-1"
        destination = "/cockroach/cockroach-data"
        read_only   = false
      }
    }
  }
}
